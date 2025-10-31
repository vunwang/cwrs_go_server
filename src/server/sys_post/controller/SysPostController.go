package controller

import (
	"cwrs_go_server/src/cwrs_common/cwrs_res"
	"cwrs_go_server/src/server/sys_post/pojo"
	"cwrs_go_server/src/server/sys_post/service"

	"github.com/gin-gonic/gin"
)

var sysPostServiceImpl = service.SysPostService{}

// @Tags 岗位管理【平台】
// @Summary 新增岗位
// @Description 新增岗位
// @Accept json
// @Produce json
// @Param Token header string true "token"
// @Param req body pojo.AddSysPostReq true "新增岗位参数"
// @Success 200 {object} cwrs_res.ResSuccess "操作成功"
// @Failure 422 {object} cwrs_res.ResError "操作失败"
// @Router /sysPost/add [post]
func AddSysPost(c *gin.Context) {
	var req pojo.AddSysPostReq
	if err := c.ShouldBindJSON(&req); err != nil {
		cwrs_res.Parameter(c, err, "参数错误！")
		return
	}
	sysPostServiceImpl.AddSysPost(c, &req)
}

// @Tags 岗位管理【平台】
// @Summary 修改岗位
// @Description 修改岗位
// @Accept json
// @Produce json
// @Param Token header string true "token"
// @Param req body pojo.EditSysPostReq true "修改岗位参数"
// @Success 200 {object} cwrs_res.ResSuccess "操作成功"
// @Failure 422 {object} cwrs_res.ResError "操作失败"
// @Router /sysPost/edit [put]
func EditSysPost(c *gin.Context) {
	var req pojo.EditSysPostReq
	if err := c.ShouldBindJSON(&req); err != nil {
		cwrs_res.Parameter(c, err, "参数错误！")
		return
	}
	sysPostServiceImpl.EditSysPost(c, &req)
}

// @Tags 岗位管理【平台】
// @Summary 删除岗位
// @Description 批量删除岗位
// @Accept json
// @Produce json
// @Param Token header string true "token"
// @Param req query pojo.DelSysPostReq true "删除岗位参数"
// @Success 200 {object} cwrs_res.ResSuccess "操作成功"
// @Failure 422 {object} cwrs_res.ResError "操作失败"
// @Router /sysPost/del [delete]
func DelSysPost(c *gin.Context) {
	var req pojo.DelSysPostReq
	if err := c.ShouldBindQuery(&req); err != nil {
		cwrs_res.Parameter(c, err, "参数错误！")
		return
	}
	sysPostServiceImpl.DelSysPost(c, &req)
}

// @Tags 岗位管理【平台】
// @Summary 查询岗位详情
// @Description 根据主键查询岗位详情
// @Accept json
// @Produce json
// @Param Token header string true "token"
// @Param req query pojo.GetSysPostDetailReq true "查询详情参数"
// @Success 200 {object} cwrs_res.ResSuccessData{data=pojo.SysPostResp} "操作成功"
// @Failure 422 {object} cwrs_res.ResError "操作失败"
// @Router /sysPost/detail [get]
func GetSysPostDetail(c *gin.Context) {
	var req pojo.GetSysPostDetailReq
	if err := c.ShouldBindQuery(&req); err != nil {
		cwrs_res.Parameter(c, err, "参数错误！")
		return
	}
	sysPostServiceImpl.GetSysPostDetail(c, &req)
}

// @Tags 岗位管理【平台】
// @Summary 分页查询岗位列表
// @Description 分页查询岗位列表
// @Accept json
// @Produce json
// @Param Token header string true "token"
// @Param req query pojo.GetSysPostListReq true "分页查询参数"
// @Success 200 {object} cwrs_res.ResSuccessDataList{data=[]pojo.SysPostResp} "操作成功"
// @Failure 422 {object} cwrs_res.ResError "操作失败"
// @Router /sysPost/list [get]
func GetSysPostList(c *gin.Context) {
	var req pojo.GetSysPostListReq
	if err := c.ShouldBindQuery(&req); err != nil {
		cwrs_res.Parameter(c, err, "参数错误！")
		return
	}
	sysPostServiceImpl.GetSysPostList(c, &req)
}
