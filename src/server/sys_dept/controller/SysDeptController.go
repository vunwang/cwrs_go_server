package controller

import (
	"cwrs_go_server/src/cwrs_common/cwrs_res"
	"cwrs_go_server/src/server/sys_dept/pojo"
	"cwrs_go_server/src/server/sys_dept/service"

	"github.com/gin-gonic/gin"
)

var sysDeptServiceImpl = service.SysDeptService{}

// @Tags 组织管理【平台】
// @Summary 新增组织
// @Description 新增组织
// @Accept json
// @Produce json
// @Param Token header string true "token"
// @Param req body pojo.AddSysDeptReq true "新增组织参数"
// @Success 200 {object} cwrs_res.ResSuccess "操作成功"
// @Failure 422 {object} cwrs_res.ResError "操作失败"
// @Router /sysDept/add [get]
func AddSysDept(c *gin.Context) {
	var req pojo.AddSysDeptReq
	if err := c.ShouldBindJSON(&req); err != nil {
		cwrs_res.Parameter(c, err, "参数错误！")
		return
	}
	sysDeptServiceImpl.AddSysDept(c, &req)
}

// @Tags 组织管理【平台】
// @Summary 修改组织
// @Description 修改组织
// @Accept json
// @Produce json
// @Param Token header string true "token"
// @Param req body pojo.EditSysDeptReq true "修改组织参数"
// @Success 200 {object} cwrs_res.ResSuccess "操作成功"
// @Failure 422 {object} cwrs_res.ResError "操作失败"
// @Router /sysDept/edit [put]
func EditSysDept(c *gin.Context) {
	var req pojo.EditSysDeptReq
	if err := c.ShouldBindJSON(&req); err != nil {
		cwrs_res.Parameter(c, err, "参数错误！")
		return
	}
	sysDeptServiceImpl.EditSysDept(c, &req)
}

// @Tags 组织管理【平台】
// @Summary 删除组织
// @Description 批量删除组织
// @Accept json
// @Produce json
// @Param Token header string true "token"
// @Param req query pojo.DelSysDeptReq true "删除组织参数"
// @Success 200 {object} cwrs_res.ResSuccess "操作成功"
// @Failure 422 {object} cwrs_res.ResError "操作失败"
// @Router /sysDept/del [delete]
func DelSysDept(c *gin.Context) {
	var req pojo.DelSysDeptReq
	if err := c.ShouldBindQuery(&req); err != nil {
		cwrs_res.Parameter(c, err, "参数错误！")
		return
	}
	sysDeptServiceImpl.DelSysDept(c, &req)
}

// @Tags 组织管理【平台】
// @Summary 查询组织详情
// @Description 根据主键查询组织详情
// @Accept json
// @Produce json
// @Param Token header string true "token"
// @Param req query pojo.GetSysDeptDetailReq true "查询详情参数"
// @Success 200 {object} cwrs_res.ResSuccessData{data=pojo.SysDept} "操作成功"
// @Failure 422 {object} cwrs_res.ResError "操作失败"
// @Router /sysDept/detail [get]
func GetSysDeptDetail(c *gin.Context) {
	var req pojo.GetSysDeptDetailReq
	if err := c.ShouldBindQuery(&req); err != nil {
		cwrs_res.Parameter(c, err, "参数错误！")
		return
	}
	sysDeptServiceImpl.GetSysDeptDetail(c, &req)
}

// @Tags 组织管理【平台】
// @Summary 查询组织树
// @Description 查询组织树（支持名称模糊查询、状态查询）
// @Accept json
// @Produce json
// @Param Token header string true "token"
// @Param req query pojo.GetSysDeptTreeReq true "查询组织树参数"
// @Success 200 {object} cwrs_res.ResSuccessData{data=[]pojo.SysDeptTreeNode} "操作成功"
// @Failure 422 {object} cwrs_res.ResError "操作失败"
// @Router /sysDept/tree [get]
func GetSysDeptTree(c *gin.Context) {
	var req pojo.GetSysDeptTreeReq
	if err := c.ShouldBindQuery(&req); err != nil {
		cwrs_res.Parameter(c, err, "参数错误！")
		return
	}
	sysDeptServiceImpl.GetSysDeptTree(c, &req)
}
