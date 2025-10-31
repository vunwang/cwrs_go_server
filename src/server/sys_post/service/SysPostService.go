package service

import (
	"cwrs_go_server/src/cwrs_common/cwrs_res"
	"cwrs_go_server/src/cwrs_utils"
	"cwrs_go_server/src/server/sys_post/dao"
	"cwrs_go_server/src/server/sys_post/pojo"
	userDao "cwrs_go_server/src/server/sys_user/dao"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

var sysPostDaoImpl = dao.SysPostDao{}
var sysUserDaoImpl = userDao.SysUserDao{}

type SysPostService struct{}

// AddSysPost 新增岗位
func (*SysPostService) AddSysPost(c *gin.Context, req *pojo.AddSysPostReq) {
	var entity pojo.SysPost
	cwrs_utils.CopyStruct(req, &entity, "json")
	entity.PostId = cwrs_utils.CreateUuid()
	entity.CreatedUserId = cwrs_utils.GetLoginUserInfo(c).UserId
	entity.CreatedTime = cwrs_utils.GetNowDateTime()
	if err := sysPostDaoImpl.AddSysPost(&entity); err != nil {
		cwrs_res.Waring(c, err, "新增岗位失败")
		return
	}
	cwrs_res.Success(c, "操作成功")
}

// EditSysPost 修改岗位
func (*SysPostService) EditSysPost(c *gin.Context, req *pojo.EditSysPostReq) {
	var entity pojo.SysPost
	cwrs_utils.CopyStruct(req, &entity, "json")
	entity.UpdatedUserId = cwrs_utils.GetLoginUserInfo(c).UserId
	entity.UpdatedTime = cwrs_utils.GetNowDateTime()
	if err := sysPostDaoImpl.EditSysPost(&entity); err != nil {
		cwrs_res.Waring(c, err, "修改岗位失败")
		return
	}
	cwrs_res.Success(c, "操作成功")
}

// DelSysPost 删除岗位
func (*SysPostService) DelSysPost(c *gin.Context, req *pojo.DelSysPostReq) {
	postIds := strings.Split(req.PostIds, ",")
	var err error
	if len(postIds) == 1 {
		if count, _ := sysUserDaoImpl.GetUserCountByPostIds(postIds[0]); count > 0 {
			cwrs_res.Waring(c, err, "岗位已被使用，无法删除")
			return
		}
		err = sysPostDaoImpl.DelSysPost(postIds[0])
		if err != nil {
			cwrs_res.Waring(c, err, "删除岗位失败")
			return
		}
		cwrs_res.Success(c, "操作成功")
	} else {
		var ids []string
		for _, postId := range postIds {
			count, _ := sysUserDaoImpl.GetUserCountByPostIds(postId)
			if count > 0 {
				ids = append(ids, postId)
				continue
			}
			err = sysPostDaoImpl.DelSysPost(postId)
		}
		if err != nil {
			cwrs_res.Waring(c, err, "删除岗位失败")
			return
		}
		cwrs_res.Success(c, fmt.Sprintf("已删除 %d 个岗位；有%d个已被使用，无法删除", len(postIds)-len(ids), len(ids)))
	}
}

// GetSysPostDetail 查询岗位详情
func (*SysPostService) GetSysPostDetail(c *gin.Context, req *pojo.GetSysPostDetailReq) {
	post, err := sysPostDaoImpl.GetSysPostById(req.PostId)
	if err != nil {
		cwrs_res.Waring(c, err, "查询岗位详情失败")
		return
	}
	cwrs_res.SuccessData(c, "操作成功", post)
}

// GetSysPostList 分页查询岗位列表
func (*SysPostService) GetSysPostList(c *gin.Context, req *pojo.GetSysPostListReq) {
	list, total, err := sysPostDaoImpl.GetSysPostList(c, req)
	if err != nil {
		cwrs_res.Waring(c, err, "查询岗位列表失败")
		return
	}
	cwrs_res.SuccessDataList(c, "操作成功", list, total)
}
