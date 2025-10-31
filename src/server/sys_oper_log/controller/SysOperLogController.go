package controller

import (
	"cwrs_go_server/src/cwrs_common/cwrs_res"
	"cwrs_go_server/src/server/sys_oper_log/pojo"
	"cwrs_go_server/src/server/sys_oper_log/service"
	"github.com/gin-gonic/gin"
)

var sysOperLogServiceImpl = service.SysOperLogService{}

// @Tags 操作日志管理【平台】
// @Summary 查询操作日志详情
// @Description 根据主键查询操作日志详情
// @Accept json
// @Produce json
// @Param Token header string true "token"
// @Param req query pojo.GetSysOperLogDetailReq true "查询详情参数"
// @Success 200 {object} cwrs_res.ResSuccessData{data=pojo.SysOperLogResp} "操作成功"
// @Failure 422 {object} cwrs_res.ResError "操作失败"
// @Router /sysOperLog/detail [get]
func GetSysOperLogDetail(c *gin.Context) {
	var req pojo.GetSysOperLogDetailReq
	if err := c.ShouldBindQuery(&req); err != nil {
		cwrs_res.Parameter(c, err, "参数错误！")
		return
	}
	sysOperLogServiceImpl.GetSysOperLogDetail(c, &req)
}

// @Tags 操作日志管理【平台】
// @Summary 分页查询操作日志列表
// @Description 分页查询操作日志列表
// @Accept json
// @Produce json
// @Param Token header string true "token"
// @Param req query pojo.GetSysOperLogListReq true "分页查询参数"
// @Success 200 {object} cwrs_res.ResSuccessDataList{data=[]pojo.SysOperLogListResp} "操作成功"
// @Failure 422 {object} cwrs_res.ResError "操作失败"
// @Router /sysOperLog/list [get]
func GetSysOperLogList(c *gin.Context) {
	var req pojo.GetSysOperLogListReq
	if err := c.ShouldBindQuery(&req); err != nil {
		cwrs_res.Parameter(c, err, "参数错误！")
		return
	}
	sysOperLogServiceImpl.GetSysOperLogList(c, &req)
}
