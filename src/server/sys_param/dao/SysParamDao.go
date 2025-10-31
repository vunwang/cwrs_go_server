package dao

import (
	"cwrs_go_server/src/cwrs_core/cwrs_gorm"
	"cwrs_go_server/src/cwrs_core/cwrs_zap_logger"
	"cwrs_go_server/src/cwrs_utils"
	"cwrs_go_server/src/server/sys_param/pojo"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var sysParamLog = cwrs_zap_logger.ZapLogger

const tableSysParam = "sys_param"
const tableSelSysParam = "sys_param sp"

type SysParamDao struct{}

// AddSysParam 添加系统参数
func (*SysParamDao) AddSysParam(item *pojo.SysParam) error {
	fields := make([]string, 0)
	if item.SysLogo != "" {
		fields = append(fields, "sys_logo")
	}
	if item.SysTitle != "" {
		fields = append(fields, "sys_title")
	}
	if item.LargeScreenTitle != "" {
		fields = append(fields, "large_screen_title")
	}
	if item.DeptId != "" {
		fields = append(fields, "dept_id")
	}
	fields = append(fields, "param_id", "created_user_id", "created_time")
	if err := cwrs_gorm.GormDb.Table(tableSysParam).Select(fields).Create(&item).Error; err != nil {
		sysParamLog.Error("AddSysParam Error", zap.Error(err))
		return err
	}
	return nil
}

// EditSysParam 修改系统参数
func (*SysParamDao) EditSysParam(item *pojo.SysParam) error {
	fields := make([]string, 0)
	if item.SysLogo != "" {
		fields = append(fields, "sys_logo")
	}
	if item.SysTitle != "" {
		fields = append(fields, "sys_title")
	}
	if item.LargeScreenTitle != "" {
		fields = append(fields, "large_screen_title")
	}
	if item.DeptId != "" {
		fields = append(fields, "dept_id")
	}
	fields = append(fields, "updated_user_id", "updated_time")
	if err := cwrs_gorm.GormDb.Table(tableSysParam).Select(fields).
		Where("param_id = ?", item.ParamId).Updates(&item).Error; err != nil {
		sysParamLog.Error("EditSysParam Error", zap.Error(err))
		return err
	}
	return nil
}

// DelSysParam 删除系统参数
func (*SysParamDao) DelSysParam(paramIds []string) error {
	if err := cwrs_gorm.GormDb.Table(tableSysParam).
		Where("param_id IN (?)", paramIds).Delete(&pojo.SysParam{}).Error; err != nil {
		sysParamLog.Error("DelSysParam Error", zap.Error(err))
		return err
	}
	return nil
}

func (*SysParamDao) GetSysParamByIds(paramIds []string) ([]pojo.SysParamResp, error) {
	var list []pojo.SysParamResp
	if err := cwrs_gorm.GormDb.Table(tableSysParam).
		Where("param_id IN (?)", paramIds).Find(&list).Error; err != nil {
		sysParamLog.Error("GetSysParamByIds Error", zap.Error(err))
		return nil, err
	}
	return list, nil
}

// GetSysParamById 查询系统参数详情
func (*SysParamDao) GetSysParamById(paramId string) (*pojo.SysParamResp, error) {
	var item pojo.SysParamResp
	if err := cwrs_gorm.GormDb.Table(tableSelSysParam).
		Select("sp.param_id", "sp.sys_logo", "sp.sys_title", "sp.large_screen_title", "sp.dept_id", "sp.created_user_id", "sp.created_time", "sd.dept_name", "su.user_name as created_user_name").
		Joins("LEFT JOIN sys_dept sd ON sp.dept_id = sd.dept_id").
		Joins("LEFT JOIN sys_user su ON sp.created_user_id = su.user_id").
		Where("sp.param_id = ?", paramId).First(&item).Error; err != nil {
		sysParamLog.Error("GetSysParamById Error", zap.Error(err))
		return nil, err
	}
	return &item, nil
}

func (*SysParamDao) GetSysParamByDeptId(deptId string) (*pojo.SysParamDeptResp, error) {
	var item pojo.SysParamDeptResp
	if err := cwrs_gorm.GormDb.Table(tableSelSysParam).
		Select("sp.param_id", "sp.sys_logo", "sp.sys_title", "sp.large_screen_title", "sp.dept_id").
		Where("sp.dept_id = ?", deptId).First(&item).Error; err != nil {
		sysParamLog.Error("GetSysParamByDeptId Error", zap.Error(err))
		return nil, err
	}
	return &item, nil
}

// GetSysParamList 分页查询系统参数列表
func (*SysParamDao) GetSysParamList(c *gin.Context, req *pojo.GetSysParamListReq) ([]pojo.SysParamResp, int64, error) {
	var list []pojo.SysParamResp
	var total int64
	db := cwrs_gorm.GormDb.Table(tableSelSysParam).Scopes(cwrs_gorm.WithDataScope(c, "sp")).
		Select("sp.param_id", "sp.sys_logo", "sp.sys_title", "sp.large_screen_title", "sp.dept_id", "sp.created_user_id", "sp.created_time", "sd.dept_name", "su.user_name as created_user_name").
		Joins("LEFT JOIN sys_dept sd ON sp.dept_id = sd.dept_id").
		Joins("LEFT JOIN sys_user su ON sp.created_user_id = su.user_id")
	if req.DeptId != "" {
		db = db.Where("sp.dept_id = ?", req.DeptId)
	}
	if req.CreatedUserName != "" {
		db = db.Where("su.user_name LIKE ?", "%"+req.CreatedUserName+"%")
	}
	//if req.StartTime != "" && req.EndTime != "" {
	//	db = db.Where("sp.created_time BETWEEN ? AND ?", req.StartTime, req.EndTime)
	//}
	db.Count(&total)
	if req.PageNum > 0 && req.PageSize > 0 {
		offset := cwrs_utils.CalcOffset(req.PageNum, req.PageSize)
		db = db.Offset(offset).Limit(req.PageSize)
	}
	if err := db.Order("sp.created_time desc").Find(&list).Error; err != nil {
		sysParamLog.Error("GetSysParamList Error", zap.Error(err))
		return nil, 0, err
	}
	return list, total, nil
}

// GetParamCountByDeptIds 获取组织下系统参数数量
func (*SysParamDao) GetParamCountByDeptIds(deptIds []string) (int64, error) {
	var count int64
	if err := cwrs_gorm.GormDb.Table(tableSysParam).
		Where("dept_id IN (?)", deptIds).
		Count(&count).Error; err != nil {
		sysParamLog.Error("GetParamCountByDeptIds Error", zap.Error(err))
		return 0, err
	}
	return count, nil
}
