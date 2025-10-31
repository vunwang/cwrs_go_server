package service

import (
	"cwrs_go_server/src/cwrs_common/cwrs_res"
	"cwrs_go_server/src/cwrs_core/cwrs_redis"
	"cwrs_go_server/src/cwrs_core/cwrs_zap_logger"
	"cwrs_go_server/src/cwrs_utils"
	"cwrs_go_server/src/server/sys_param/dao"
	"cwrs_go_server/src/server/sys_param/pojo"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"strings"
)

var sysParamDaoImpl = dao.SysParamDao{}

type SysParamService struct{}

// AddSysParam 新增系统参数
func (*SysParamService) AddSysParam(c *gin.Context, req *pojo.AddSysParamReq) {
	var entity pojo.SysParam
	cwrs_utils.CopyStruct(req, &entity, "json")
	entity.ParamId = cwrs_utils.CreateUuid()
	entity.CreatedUserId = cwrs_utils.GetLoginUserInfo(c).UserId
	entity.CreatedTime = cwrs_utils.GetNowDateTime()
	if err := sysParamDaoImpl.AddSysParam(&entity); err != nil {
		if strings.Contains(err.Error(), "only_dept_id") {
			cwrs_res.Waring(c, err, "该组织下已存在相关参数，不可重复添加")
			return
		}
		cwrs_res.Waring(c, err, "新增系统参数失败")
		return
	}
	jsonData, err := json.Marshal(req)
	addRedisSysParam(c, entity.DeptId, jsonData, err)
	cwrs_res.Success(c, "操作成功")
}

// EditSysParam 修改系统参数
func (*SysParamService) EditSysParam(c *gin.Context, req *pojo.EditSysParamReq) {
	var entity pojo.SysParam
	cwrs_utils.CopyStruct(req, &entity, "json")
	entity.UpdatedUserId = cwrs_utils.GetLoginUserInfo(c).UserId
	entity.UpdatedTime = cwrs_utils.GetNowDateTime()
	if err := sysParamDaoImpl.EditSysParam(&entity); err != nil {
		if strings.Contains(err.Error(), "only_dept_id") {
			cwrs_res.Waring(c, err, "该组织下已存在相关参数，不可重复添加")
			return
		}
		cwrs_res.Waring(c, err, "修改系统参数失败")
		return
	}
	jsonData, err := json.Marshal(req)
	addRedisSysParam(c, entity.DeptId, jsonData, err)
	// 判断是否删除oss中图片
	resObj, err := sysParamDaoImpl.GetSysParamById(entity.ParamId)
	if err != nil {
		cwrs_zap_logger.Error("EditSysParam-删除oss文件时获取参数失败", zap.Error(err))
	} else {
		if resObj.SysLogo != entity.SysLogo {
			err = cwrs_utils.AliyunOssDeleteObject(resObj.SysLogo)
			if err != nil {
				cwrs_zap_logger.Error("EditSysParam-删除oss文件失败", zap.Error(err))
			}
		}
	}
	cwrs_res.Success(c, "操作成功")
}

// 添加数据到redis
func addRedisSysParam(c *gin.Context, deptId string, jsonData []byte, err error) {
	if err != nil {
		cwrs_zap_logger.Error("addRedisSysParam-json序列化失败", zap.Error(err))
	}
	redisKey := fmt.Sprintf("%s%s", cwrs_redis.KEY_SYS_PARAM, deptId)
	if err = cwrs_redis.GlobalRedis.Set(c.Request.Context(), redisKey, jsonData, 0).Err(); err != nil {
		cwrs_zap_logger.Error("addRedisSysParam-保存数据到redis失败", zap.Error(err))
	}
}

// DelSysParam 删除系统参数
func (*SysParamService) DelSysParam(c *gin.Context, req *pojo.DelSysParamReq) {
	paramIds := strings.Split(req.ParamIds, ",")
	if err := sysParamDaoImpl.DelSysParam(paramIds); err != nil {
		cwrs_res.Waring(c, err, "删除系统参数失败")
		return
	}
	list, err := sysParamDaoImpl.GetSysParamByIds(paramIds)
	if err != nil {
		cwrs_zap_logger.Error("DelSysParam-查询系统参数失败", zap.Error(err))
	} else {
		redisKeys := make([]string, 0)
		for _, item := range list {
			redisKeys = append(redisKeys, fmt.Sprintf("%s%s", cwrs_redis.KEY_SYS_PARAM, item.DeptId))
			//删除oss文件
			if item.SysLogo != "" {
				err = cwrs_utils.AliyunOssDeleteObject(item.SysLogo)
				if err != nil {
					cwrs_zap_logger.Error("DelSysParam-删除oss文件失败", zap.Error(err))
				}
			}
		}
		cwrs_redis.GlobalRedis.Del(c.Request.Context(), redisKeys...)
	}
	cwrs_res.Success(c, "操作成功")
}

// GetSysParamDetail 查询系统参数详情
func (*SysParamService) GetSysParamDetail(c *gin.Context, req *pojo.GetSysParamDetailReq) {
	item, err := sysParamDaoImpl.GetSysParamById(req.ParamId)
	if err != nil {
		cwrs_res.Waring(c, err, "查询系统参数详情失败")
		return
	}
	cwrs_res.SuccessData(c, "操作成功", item)
}

// GetSysParamDept 根据组织id查询系统参数
func (*SysParamService) GetSysParamByDeptId(c *gin.Context) {
	var item *pojo.SysParamDeptResp
	deptId := cwrs_utils.GetLoginUserInfo(c).DeptId
	jsonData, err := cwrs_redis.GlobalRedis.Get(c.Request.Context(), fmt.Sprintf("%s%s", cwrs_redis.KEY_SYS_PARAM, deptId)).Result()
	if err != nil {
		cwrs_zap_logger.Error("GetSysParamByDeptId-查询redis失败", zap.Error(err))
	}
	if jsonData != "" {
		err = json.Unmarshal([]byte(jsonData), &item)
		if err != nil {
			cwrs_res.Waring(c, err, "根据组织id查询系统参数失败")
			return
		}
	} else {
		jsonData, err = cwrs_redis.GlobalRedis.Get(c.Request.Context(), fmt.Sprintf("%s%s", cwrs_redis.KEY_SYS_PARAM, "all")).Result()
		if err != nil {
			cwrs_zap_logger.Error("GetSysParamByDeptId-查询redis失败", zap.Error(err))
		}
		if jsonData != "" {
			err = json.Unmarshal([]byte(jsonData), &item)
			if err != nil {
				cwrs_res.Waring(c, err, "根据组织id查询系统参数失败")
				return
			}
		} else {
			//redis数据没获取到 从数据库获取
			item, err = sysParamDaoImpl.GetSysParamByDeptId(deptId)
			if err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					item, err = sysParamDaoImpl.GetSysParamByDeptId("all")
					if err != nil {
						cwrs_res.Waring(c, err, "根据组织id查询系统参数失败")
						return
					}
				} else {
					cwrs_res.Waring(c, err, "根据组织id查询系统参数失败")
					return
				}
			}
			jsonDataRes, errRes := json.Marshal(item)
			addRedisSysParam(c, deptId, jsonDataRes, errRes)
		}
	}
	cwrs_res.SuccessData(c, "操作成功", item)
}

// GetSysParamList 分页查询系统参数列表
func (*SysParamService) GetSysParamList(c *gin.Context, req *pojo.GetSysParamListReq) {
	list, total, err := sysParamDaoImpl.GetSysParamList(c, req)
	if err != nil {
		cwrs_res.Waring(c, err, "查询系统参数列表失败")
		return
	}
	cwrs_res.SuccessDataList(c, "操作成功", list, total)
}
