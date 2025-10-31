package controller

import (
	"cwrs_go_server/src/cwrs_common/cwrs_res"
	"cwrs_go_server/src/server/sys_param/pojo"
	"cwrs_go_server/src/server/sys_param/service"
	"github.com/gin-gonic/gin"
)

var sysParamServiceImpl = service.SysParamService{}

// @Tags 系统参数管理【平台】
// @Summary 新增系统参数
// @Description 新增系统参数
// @Accept json
// @Produce json
// @Param Token header string true "token"
// @Param req body pojo.AddSysParamReq true "新增参数"
// @Success 200 {object} cwrs_res.ResSuccess "操作成功"
// @Failure 422 {object} cwrs_res.ResError "操作失败"
// @Router /sysParam/add [post]
func AddSysParam(c *gin.Context) {
	var req pojo.AddSysParamReq
	if err := c.ShouldBindJSON(&req); err != nil {
		cwrs_res.Parameter(c, err, "参数错误！")
		return
	}
	sysParamServiceImpl.AddSysParam(c, &req)
}

// @Tags 系统参数管理【平台】
// @Summary 修改系统参数
// @Description 修改系统参数
// @Accept json
// @Produce json
// @Param Token header string true "token"
// @Param req body pojo.EditSysParamReq true "修改参数"
// @Success 200 {object} cwrs_res.ResSuccess "操作成功"
// @Failure 422 {object} cwrs_res.ResError "操作失败"
// @Router /sysParam/edit [put]
func EditSysParam(c *gin.Context) {
	var req pojo.EditSysParamReq
	if err := c.ShouldBindJSON(&req); err != nil {
		cwrs_res.Parameter(c, err, "参数错误！")
		return
	}
	sysParamServiceImpl.EditSysParam(c, &req)
}

// @Tags 系统参数管理【平台】
// @Summary 删除系统参数 所属组织值为all时，隐藏删除按钮
// @Description 删除系统参数 所属组织值为all时，隐藏删除按钮
// @Accept json
// @Produce json
// @Param Token header string true "token"
// @Param req query pojo.DelSysParamReq true "删除参数"
// @Success 200 {object} cwrs_res.ResSuccess "操作成功"
// @Failure 422 {object} cwrs_res.ResError "操作失败"
// @Router /sysParam/del [delete]
func DelSysParam(c *gin.Context) {
	var req pojo.DelSysParamReq
	if err := c.ShouldBindQuery(&req); err != nil {
		cwrs_res.Parameter(c, err, "参数错误！")
		return
	}
	sysParamServiceImpl.DelSysParam(c, &req)
}

// @Tags 系统参数管理【平台】
// @Summary 获取当前登录用户所属组织的系统参数
// @Description 获取当前登录用户所属组织的系统参数
// @Accept json
// @Produce json
// @Param Token header string true "token"
// @Success 200 {object} cwrs_res.ResSuccessData{data=pojo.SysParamDeptResp} "操作成功"
// @Failure 422 {object} cwrs_res.ResError "操作失败"
// @Router /sysParam/dept [get]
func GetSysParamByDeptId(c *gin.Context) {
	sysParamServiceImpl.GetSysParamByDeptId(c)
}

// @Tags 系统参数管理【平台】
// @Summary 查询系统参数详情
// @Description 根据主键查询系统参数详情
// @Accept json
// @Produce json
// @Param Token header string true "token"
// @Param req query pojo.GetSysParamDetailReq true "查询详情参数"
// @Success 200 {object} cwrs_res.ResSuccessData{data=pojo.SysParamResp} "操作成功"
// @Failure 422 {object} cwrs_res.ResError "操作失败"
// @Router /sysParam/detail [get]
func GetSysParamDetail(c *gin.Context) {
	var req pojo.GetSysParamDetailReq
	if err := c.ShouldBindQuery(&req); err != nil {
		cwrs_res.Parameter(c, err, "参数错误！")
		return
	}
	sysParamServiceImpl.GetSysParamDetail(c, &req)
}

// @Tags 系统参数管理【平台】
// @Summary 分页查询系统参数列表
// @Description 分页查询系统参数列表
// @Accept json
// @Produce json
// @Param Token header string true "token"
// @Param req query pojo.GetSysParamListReq true "分页查询参数"
// @Success 200 {object} cwrs_res.ResSuccessDataList{data=[]pojo.SysParamResp} "操作成功"
// @Failure 422 {object} cwrs_res.ResError "操作失败"
// @Router /sysParam/list [get]
func GetSysParamList(c *gin.Context) {
	var req pojo.GetSysParamListReq
	if err := c.ShouldBindQuery(&req); err != nil {
		cwrs_res.Parameter(c, err, "参数错误！")
		return
	}
	sysParamServiceImpl.GetSysParamList(c, &req)
}
