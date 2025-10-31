package controller

import (
	"cwrs_go_server/src/cwrs_common/cwrs_res"
	"cwrs_go_server/src/server/sys_user/pojo"
	"cwrs_go_server/src/server/sys_user/service"
	"github.com/gin-gonic/gin"
)

var sysUserServiceImpl = service.SysUserService{}

// @Tags 用户管理【平台】
// @Summary 新增用户
// @Description 新增用户
// @Accept json
// @Produce json
// @Param Token header string true "token"
// @Param req body pojo.AddSysUserReq true "新增参数"
// @Success 200 {object} cwrs_res.ResSuccess "操作成功"
// @Failure 422 {object} cwrs_res.ResError "操作失败"
// @Router /sysUser/add [post]
func AddSysUser(c *gin.Context) {
	var req pojo.AddSysUserReq
	if err := c.ShouldBindJSON(&req); err != nil {
		cwrs_res.Parameter(c, err, "参数错误！")
		return
	}
	sysUserServiceImpl.AddSysUser(c, &req)
}

// @Tags 用户管理【平台】
// @Summary 重置用户密码
// @Description 重置用户密码
// @Accept json
// @Produce json
// @Param Token header string true "token"
// @Param req body pojo.GetSysUserDetailReq true "重置密码入参"
// @Success 200 {object} cwrs_res.ResSuccess "操作成功"
// @Failure 422 {object} cwrs_res.ResError "操作失败"
// @Router /sysUser/resetPassword [put]
func ResetUserPassword(c *gin.Context) {
	var req pojo.GetSysUserDetailReq
	if err := c.ShouldBindJSON(&req); err != nil {
		cwrs_res.Parameter(c, err, "参数错误！")
		return
	}
	sysUserServiceImpl.ResetUserPassword(c, &req)
}

// @Tags 用户管理【平台】
// @Summary 用户自己修改密码
// @Description 用户自己修改密码
// @Accept json
// @Produce json
// @Param Token header string true "token"
// @Param req body pojo.EditUserPwdReq true "修改密码参数"
// @Success 200 {object} cwrs_res.ResSuccess "操作成功"
// @Failure 422 {object} cwrs_res.ResError "操作失败"
// @Router /sysUser/editPassword [put]
func EditUserPassword(c *gin.Context) {
	var req pojo.EditUserPwdReq
	if err := c.ShouldBindJSON(&req); err != nil {
		cwrs_res.Parameter(c, err, "参数错误！")
		return
	}
	sysUserServiceImpl.EditUserPassword(c, &req)
}

// @Tags 用户管理【平台】
// @Summary 修改用户
// @Description 修改用户
// @Accept json
// @Produce json
// @Param Token header string true "token"
// @Param req body pojo.EditSysUserReq true "修改参数"
// @Success 200 {object} cwrs_res.ResSuccess "操作成功"
// @Failure 422 {object} cwrs_res.ResError "操作失败"
// @Router /sysUser/edit [put]
func EditSysUser(c *gin.Context) {
	var req pojo.EditSysUserReq
	if err := c.ShouldBindJSON(&req); err != nil {
		cwrs_res.Parameter(c, err, "参数错误！")
		return
	}
	sysUserServiceImpl.EditSysUser(c, &req)
}

// @Tags 用户管理【平台】
// @Summary 删除用户
// @Description 删除用户
// @Accept json
// @Produce json
// @Param Token header string true "token"
// @Param req query pojo.DelSysUserReq true "删除参数"
// @Success 200 {object} cwrs_res.ResSuccess "操作成功"
// @Failure 422 {object} cwrs_res.ResError "操作失败"
// @Router /sysUser/del [delete]
func DelSysUser(c *gin.Context) {
	var req pojo.DelSysUserReq
	if err := c.ShouldBindQuery(&req); err != nil {
		cwrs_res.Parameter(c, err, "参数错误！")
		return
	}
	sysUserServiceImpl.DelSysUser(c, &req)
}

// @Tags 用户管理【平台】
// @Summary 查询用户详情
// @Description 根据主键查询用户详情
// @Accept json
// @Produce json
// @Param Token header string true "token"
// @Param req query pojo.GetSysUserDetailReq true "查询详情参数"
// @Success 200 {object} cwrs_res.ResSuccessData{data=pojo.SysUserResp} "操作成功"
// @Failure 422 {object} cwrs_res.ResError "操作失败"
// @Router /sysUser/detail [get]
func GetSysUserDetail(c *gin.Context) {
	var req pojo.GetSysUserDetailReq
	if err := c.ShouldBindQuery(&req); err != nil {
		cwrs_res.Parameter(c, err, "参数错误！")
		return
	}
	sysUserServiceImpl.GetSysUserDetail(c, &req)
}

// @Tags 用户管理【平台】
// @Summary 分页查询用户列表
// @Description 分页查询用户列表
// @Accept json
// @Produce json
// @Param Token header string true "token"
// @Param req query pojo.GetSysUserListReq true "分页查询参数"
// @Success 200 {object} cwrs_res.ResSuccessDataList{data=[]pojo.SysUserListResp} "操作成功"
// @Failure 422 {object} cwrs_res.ResError "操作失败"
// @Router /sysUser/list [get]
func GetSysUserList(c *gin.Context) {
	var req pojo.GetSysUserListReq
	if err := c.ShouldBindQuery(&req); err != nil {
		cwrs_res.Parameter(c, err, "参数错误！")
		return
	}
	sysUserServiceImpl.GetSysUserList(c, &req)
}
