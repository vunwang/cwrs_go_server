package controller

import (
	"cwrs_go_server/src/cwrs_common/cwrs_res"
	"cwrs_go_server/src/server/sys_task/pojo"
	"cwrs_go_server/src/server/sys_task/service"
	"github.com/gin-gonic/gin"
)

var sysTaskServiceImpl = service.SysTaskService{}
var taskServiceImpl = service.TaskRegistry{}

// @Tags 定时任务管理【平台】
// @Summary 打开定时任务
// @Description 打开定时任务
// @Accept json
// @Produce json
// @Param Token header string true "token"
// @Param req query pojo.StartAndStopTaskReq true "打开任务参数"
// @Success 200 {object} cwrs_res.ResSuccess "操作成功"
// @Failure 422 {object} cwrs_res.ResError "操作失败"
// @Router /sysTask/start [get]
func StartTask(c *gin.Context) {
	var req pojo.StartAndStopTaskReq
	if err := c.ShouldBindQuery(&req); err != nil {
		cwrs_res.Parameter(c, err, "参数错误！")
		return
	}
	taskServiceImpl.StartTask(c, &req)
}

// @Tags 定时任务管理【平台】
// @Summary 关闭定时任务
// @Description 关闭定时任务
// @Accept json
// @Produce json
// @Param Token header string true "token"
// @Param req query pojo.StartAndStopTaskReq true "关闭任务参数"
// @Success 200 {object} cwrs_res.ResSuccess "操作成功"
// @Failure 422 {object} cwrs_res.ResError "操作失败"
// @Router /sysTask/stop [get]
func StopTask(c *gin.Context) {
	var req pojo.StartAndStopTaskReq
	if err := c.ShouldBindQuery(&req); err != nil {
		cwrs_res.Parameter(c, err, "参数错误！")
		return
	}
	taskServiceImpl.StopTask(c, &req)
}

// @Tags 定时任务管理【平台】
// @Summary 新增定时任务
// @Description 新增定时任务
// @Accept json
// @Produce json
// @Param Token header string true "token"
// @Param req body pojo.AddSysTaskReq true "新增参数"
// @Success 200 {object} cwrs_res.ResSuccess "操作成功"
// @Failure 422 {object} cwrs_res.ResError "操作失败"
// @Router /sysTask/add [post]
func AddSysTask(c *gin.Context) {
	var req pojo.AddSysTaskReq
	if err := c.ShouldBindJSON(&req); err != nil {
		cwrs_res.Parameter(c, err, "参数错误！")
		return
	}
	sysTaskServiceImpl.AddSysTask(c, &req)
}

// @Tags 定时任务管理【平台】
// @Summary 修改定时任务 只能修改停止状态的任务
// @Description 修改定时任务 只能修改停止状态的任务
// @Accept json
// @Produce json
// @Param Token header string true "token"
// @Param req body pojo.EditSysTaskReq true "修改参数"
// @Success 200 {object} cwrs_res.ResSuccess "操作成功"
// @Failure 422 {object} cwrs_res.ResError "操作失败"
// @Router /sysTask/edit [put]
func EditSysTask(c *gin.Context) {
	var req pojo.EditSysTaskReq
	if err := c.ShouldBindJSON(&req); err != nil {
		cwrs_res.Parameter(c, err, "参数错误！")
		return
	}
	sysTaskServiceImpl.EditSysTask(c, &req)
}

// @Tags 定时任务管理【平台】
// @Summary 删除定时任务 只能删除停止状态的任务
// @Description 删除定时任务 只能删除停止状态的任务
// @Accept json
// @Produce json
// @Param Token header string true "token"
// @Param req query pojo.DelSysTaskReq true "删除参数"
// @Success 200 {object} cwrs_res.ResSuccess "操作成功"
// @Failure 422 {object} cwrs_res.ResError "操作失败"
// @Router /sysTask/del [delete]
func DelSysTask(c *gin.Context) {
	var req pojo.DelSysTaskReq
	if err := c.ShouldBindQuery(&req); err != nil {
		cwrs_res.Parameter(c, err, "参数错误！")
		return
	}
	sysTaskServiceImpl.DelSysTask(c, &req)
}

// @Tags 定时任务管理【平台】
// @Summary 查询定时任务详情
// @Description 根据主键查询定时任务详情
// @Accept json
// @Produce json
// @Param Token header string true "token"
// @Param req query pojo.GetSysTaskDetailReq true "查询详情参数"
// @Success 200 {object} cwrs_res.ResSuccessData{data=pojo.SysTaskResp} "操作成功"
// @Failure 422 {object} cwrs_res.ResError "操作失败"
// @Router /sysTask/detail [get]
func GetSysTaskDetail(c *gin.Context) {
	var req pojo.GetSysTaskDetailReq
	if err := c.ShouldBindQuery(&req); err != nil {
		cwrs_res.Parameter(c, err, "参数错误！")
		return
	}
	sysTaskServiceImpl.GetSysTaskDetail(c, &req)
}

// @Tags 定时任务管理【平台】
// @Summary 分页查询定时任务列表
// @Description 分页查询定时任务列表
// @Accept json
// @Produce json
// @Param Token header string true "token"
// @Param req query pojo.GetSysTaskListReq true "分页查询参数"
// @Success 200 {object} cwrs_res.ResSuccessDataList{data=[]pojo.SysTaskResp} "操作成功"
// @Failure 422 {object} cwrs_res.ResError "操作失败"
// @Router /sysTask/list [get]
func GetSysTaskList(c *gin.Context) {
	var req pojo.GetSysTaskListReq
	if err := c.ShouldBindQuery(&req); err != nil {
		cwrs_res.Parameter(c, err, "参数错误！")
		return
	}
	sysTaskServiceImpl.GetSysTaskList(c, &req)
}
