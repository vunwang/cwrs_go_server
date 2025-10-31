package service

import (
	"cwrs_go_server/src/cwrs_common/cwrs_res"
	"cwrs_go_server/src/cwrs_utils"
	"cwrs_go_server/src/server/sys_task/dao"
	"cwrs_go_server/src/server/sys_task/pojo"
	"github.com/gin-gonic/gin"
	"strings"
)

var sysTaskDaoImpl = dao.SysTaskDao{}

type SysTaskService struct{}

// AddSysTask 新增定时任务
func (*SysTaskService) AddSysTask(c *gin.Context, req *pojo.AddSysTaskReq) {
	var entity pojo.SysTask
	cwrs_utils.CopyStruct(req, &entity, "json")
	if !strings.Contains(entity.CronName, ",") {
		cwrs_res.Waring(c, nil, "执行时间格式错误")
		return
	}
	cronNames := strings.Split(req.CronName, ",")
	entity.TaskId = cwrs_utils.CreateUuid()
	entity.TaskStatus = "2"
	entity.CronName = cronNames[0]
	entity.CronExpr = cronNames[1]
	entity.CreatedUserId = cwrs_utils.GetLoginUserInfo(c).UserId
	entity.CreatedTime = cwrs_utils.GetNowDateTime()
	if err := sysTaskDaoImpl.AddSysTask(&entity); err != nil {
		if strings.Contains(err.Error(), "uniq_task_name") {
			cwrs_res.Waring(c, err, "新增定时任务失败，任务名称已存在")
		}
		cwrs_res.Waring(c, err, "新增定时任务失败")
		return
	}
	cwrs_res.Success(c, "操作成功")
}

// EditSysTask 修改定时任务
func (*SysTaskService) EditSysTask(c *gin.Context, req *pojo.EditSysTaskReq) {
	var entity pojo.SysTask
	cwrs_utils.CopyStruct(req, &entity, "json")
	if !strings.Contains(entity.CronName, ",") {
		cwrs_res.Waring(c, nil, "执行时间格式错误")
		return
	}
	cronNames := strings.Split(req.CronName, ",")
	entity.CronName = cronNames[0]
	entity.CronExpr = cronNames[1]
	entity.UpdatedUserId = cwrs_utils.GetLoginUserInfo(c).UserId
	entity.UpdatedTime = cwrs_utils.GetNowDateTime()
	if err := sysTaskDaoImpl.EditSysTask(&entity); err != nil {
		if strings.Contains(err.Error(), "uniq_task_name") {
			cwrs_res.Waring(c, err, "修改定时任务失败，任务名称已存在")
		}
		cwrs_res.Waring(c, err, "修改定时任务失败")
		return
	}
	cwrs_res.Success(c, "操作成功")
}

// DelSysTask 删除定时任务
func (*SysTaskService) DelSysTask(c *gin.Context, req *pojo.DelSysTaskReq) {
	taskIds := strings.Split(req.TaskIds, ",")
	if err := sysTaskDaoImpl.DelSysTask(taskIds); err != nil {
		cwrs_res.Waring(c, err, "删除定时任务失败")
		return
	}
	cwrs_res.Success(c, "操作成功")
}

// GetSysTaskDetail 查询定时任务详情
func (*SysTaskService) GetSysTaskDetail(c *gin.Context, req *pojo.GetSysTaskDetailReq) {
	item, err := sysTaskDaoImpl.GetSysTaskById(req.TaskId)
	if err != nil {
		cwrs_res.Waring(c, err, "查询定时任务详情失败")
		return
	}
	cwrs_res.SuccessData(c, "操作成功", item)
}

// GetSysTaskList 分页查询定时任务列表
func (*SysTaskService) GetSysTaskList(c *gin.Context, req *pojo.GetSysTaskListReq) {
	list, total, err := sysTaskDaoImpl.GetSysTaskList(req)
	if err != nil {
		cwrs_res.Waring(c, err, "查询定时任务列表失败")
		return
	}
	cwrs_res.SuccessDataList(c, "操作成功", list, total)
}
