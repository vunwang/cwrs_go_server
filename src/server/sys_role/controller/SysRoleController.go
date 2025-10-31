package controller

import (
	"cwrs_go_server/src/cwrs_common/cwrs_res"
	"cwrs_go_server/src/server/sys_role/pojo"
	"cwrs_go_server/src/server/sys_role/service"

	"github.com/gin-gonic/gin"
)

var sysRoleServiceImpl = service.SysRoleService{}

// @Tags 角色管理【平台】
// @Summary 新增角色
// @Description 新增角色
// @Accept json
// @Produce json
// @Param Token header string true "token"
// @Param req body pojo.AddSysRoleReq true "新增角色参数"
// @Success 200 {object} cwrs_res.ResSuccess "操作成功"
// @Failure 422 {object} cwrs_res.ResError "操作失败"
// @Router /sysRole/add [post]
func AddSysRole(c *gin.Context) {
	var req pojo.AddSysRoleReq
	if err := c.ShouldBindJSON(&req); err != nil {
		cwrs_res.Parameter(c, err, "参数错误！")
		return
	}
	sysRoleServiceImpl.AddSysRole(c, &req)
}

// @Tags 角色管理【平台】
// @Summary 修改角色
// @Description 修改角色
// @Accept json
// @Produce json
// @Param Token header string true "token"
// @Param req body pojo.EditSysRoleReq true "修改角色参数"
// @Success 200 {object} cwrs_res.ResSuccess "操作成功"
// @Failure 422 {object} cwrs_res.ResError "操作失败"
// @Router /sysRole/edit [put]
func EditSysRole(c *gin.Context) {
	var req pojo.EditSysRoleReq
	if err := c.ShouldBindJSON(&req); err != nil {
		cwrs_res.Parameter(c, err, "参数错误！")
		return
	}
	sysRoleServiceImpl.EditSysRole(c, &req)
}

// @Tags 角色管理【平台】
// @Summary 删除角色
// @Description 批量删除角色
// @Accept json
// @Produce json
// @Param Token header string true "token"
// @Param req query pojo.DelSysRoleReq true "删除角色参数"
// @Success 200 {object} cwrs_res.ResSuccess "操作成功"
// @Failure 422 {object} cwrs_res.ResError "操作失败"
// @Router /sysRole/del [delete]
func DelSysRole(c *gin.Context) {
	var req pojo.DelSysRoleReq
	if err := c.ShouldBindQuery(&req); err != nil {
		cwrs_res.Parameter(c, err, "参数错误！")
		return
	}
	sysRoleServiceImpl.DelSysRole(c, &req)
}

// @Tags 角色管理【平台】
// @Summary 查询角色详情
// @Description 根据主键查询角色详情
// @Accept json
// @Produce json
// @Param Token header string true "token"
// @Param req query pojo.GetSysRoleDetailReq true "查询详情参数"
// @Success 200 {object} cwrs_res.ResSuccessData{data=pojo.SysRoleDetailResp} "操作成功"
// @Failure 422 {object} cwrs_res.ResError "操作失败"
// @Router /sysRole/detail [get]
func GetSysRoleDetail(c *gin.Context) {
	var req pojo.GetSysRoleDetailReq
	if err := c.ShouldBindQuery(&req); err != nil {
		cwrs_res.Parameter(c, err, "参数错误！")
		return
	}
	sysRoleServiceImpl.GetSysRoleDetail(c, &req)
}

// @Tags 角色管理【平台】
// @Summary 分页查询角色列表
// @Description 分页查询角色列表
// @Accept json
// @Produce json
// @Param Token header string true "token"
// @Param req query pojo.GetSysRoleListReq true "分页查询参数"
// @Success 200 {object} cwrs_res.ResSuccessDataList{data=[]pojo.SysRoleResp} "操作成功"
// @Failure 422 {object} cwrs_res.ResError "操作失败"
// @Router /sysRole/list [get]
func GetSysRoleList(c *gin.Context) {
	var req pojo.GetSysRoleListReq
	if err := c.ShouldBindQuery(&req); err != nil {
		cwrs_res.Parameter(c, err, "参数错误！")
		return
	}
	sysRoleServiceImpl.GetSysRoleList(c, &req)
}
