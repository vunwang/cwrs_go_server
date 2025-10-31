package dao

import (
	"cwrs_go_server/src/cwrs_core/cwrs_gorm"
	"cwrs_go_server/src/cwrs_core/cwrs_zap_logger"
	"cwrs_go_server/src/cwrs_utils"
	"cwrs_go_server/src/server/sys_task/pojo"
	"go.uber.org/zap"
)

var sysTaskLog = cwrs_zap_logger.ZapLogger

const tableSysTask = "sys_task"
const tableSelSysTask = "sys_task st"

type SysTaskDao struct{}

// AddSysTask 添加定时任务
func (*SysTaskDao) AddSysTask(item *pojo.SysTask) error {
	fields := make([]string, 0)
	if item.TaskName != "" {
		fields = append(fields, "task_name")
	}
	if item.CronName != "" {
		fields = append(fields, "cron_name")
	}
	if item.CronExpr != "" {
		fields = append(fields, "cron_expr")
	}
	if item.FuncName != "" {
		fields = append(fields, "func_name")
	}
	if item.TaskParams != nil {
		fields = append(fields, "task_params")
	}
	if item.TaskStatus != "" {
		fields = append(fields, "task_status")
	}
	if item.Desc != "" {
		fields = append(fields, "desc")
	}
	fields = append(fields, "task_id", "created_user_id", "created_time")
	if err := cwrs_gorm.GormDb.Table(tableSysTask).Select(fields).Create(&item).Error; err != nil {
		sysTaskLog.Error("AddSysTask Error", zap.Error(err))
		return err
	}
	return nil
}

// EditSysTask 修改定时任务
func (*SysTaskDao) EditSysTask(item *pojo.SysTask) error {
	fields := make([]string, 0)
	if item.TaskName != "" {
		fields = append(fields, "task_name")
	}
	if item.CronName != "" {
		fields = append(fields, "cron_name")
	}
	if item.CronExpr != "" {
		fields = append(fields, "cron_expr")
	}
	if item.FuncName != "" {
		fields = append(fields, "func_name")
	}
	if item.TaskParams != nil {
		fields = append(fields, "task_params")
	}
	if item.TaskStatus != "" {
		fields = append(fields, "task_status")
	}
	if item.Desc != "" {
		fields = append(fields, "desc")
	}
	fields = append(fields, "updated_user_id", "updated_time")
	if err := cwrs_gorm.GormDb.Table(tableSysTask).Select(fields).
		Where("task_status = ? AND task_id = ?", "2", item.TaskId).Updates(&item).Error; err != nil {
		sysTaskLog.Error("EditSysTask Error", zap.Error(err))
		return err
	}
	return nil
}

// EditSysTaskStatus 修改定时任务状态
func (*SysTaskDao) EditSysTaskStatus(taskStatus, lastRunTime string, taskIds []string) error {
	if err := cwrs_gorm.GormDb.Table(tableSysTask).Select("task_status", "last_run_time").
		Where("task_id IN (?)", taskIds).Updates(map[string]interface{}{
		"task_status":   taskStatus,
		"last_run_time": lastRunTime,
	}).Error; err != nil {
		sysTaskLog.Error("EditSysTaskStatus Error", zap.Error(err))
		return err
	}
	return nil
}

// DelSysTask 删除定时任务
func (*SysTaskDao) DelSysTask(taskIds []string) error {
	for _, taskId := range taskIds {
		err := cwrs_gorm.GormDb.Table(tableSysTask).
			Where("task_status = ? AND task_id = ?", "2", taskId).Delete(&pojo.SysTask{}).Error
		if err != nil {
			sysTaskLog.Error("DelSysTask Error", zap.Error(err))
			return err
		}
	}
	return nil
}

// GetSysTaskById 查询定时任务详情
func (*SysTaskDao) GetSysTaskById(taskId string) (*pojo.SysTaskResp, error) {
	var item pojo.SysTaskResp
	if err := cwrs_gorm.GormDb.Table(tableSelSysTask).
		Select("st.task_id", "st.task_name", "st.cron_name", "st.cron_expr", "st.func_name", "st.task_params", "st.task_status", "st.desc", "st.last_run_time", "st.created_user_id", "st.created_time", "su.user_name as created_user_name").
		Joins("LEFT JOIN sys_user su ON st.created_user_id = su.user_id").
		Where("st.task_id = ?", taskId).First(&item).Error; err != nil {
		sysTaskLog.Error("GetSysTaskById Error", zap.Error(err))
		return nil, err
	}
	return &item, nil
}

// GetSysTaskListByIds 批量查询定时任务详情
func (*SysTaskDao) GetSysTaskListByIds(taskIds []string) ([]pojo.SysTaskResp, error) {
	var list []pojo.SysTaskResp
	if err := cwrs_gorm.GormDb.Table(tableSelSysTask).
		Select("st.task_id", "st.task_name", "st.cron_name", "st.cron_expr", "st.func_name", "st.task_params", "st.task_status", "st.desc", "st.last_run_time", "st.created_user_id", "st.created_time").
		Where("st.task_id IN (?)", taskIds).
		Scan(&list).Error; err != nil {
		sysTaskLog.Error("GetSysTaskListByIds Error", zap.Error(err))
		return nil, err
	}
	return list, nil
}

// GetSysTaskListByStatus 根据状态查询定时任务列表
func (*SysTaskDao) GetSysTaskListByStatus(taskStatus string) ([]pojo.SysTaskResp, error) {
	var list []pojo.SysTaskResp
	if err := cwrs_gorm.GormDb.Table(tableSelSysTask).
		Select("st.task_id", "st.task_name", "st.cron_name", "st.cron_expr", "st.func_name", "st.task_params", "st.task_status", "st.desc", "st.last_run_time", "st.created_user_id", "st.created_time").
		Where("st.task_status = ?", taskStatus).
		Find(&list).Error; err != nil {
		sysTaskLog.Error("GetSysTaskListByStatus Error", zap.Error(err))
		return nil, err
	}
	return list, nil
}

// GetSysTaskList 分页查询定时任务列表
func (*SysTaskDao) GetSysTaskList(req *pojo.GetSysTaskListReq) ([]pojo.SysTaskResp, int64, error) {
	var list []pojo.SysTaskResp
	var total int64
	db := cwrs_gorm.GormDb.Table(tableSelSysTask).
		Select("st.task_id", "st.task_name", "st.cron_name", "st.cron_expr", "st.func_name", "st.task_params", "st.task_status", "st.desc", "st.last_run_time", "st.created_user_id", "st.created_time", "su.user_name as created_user_name").
		Joins("LEFT JOIN sys_user su ON st.created_user_id = su.user_id")
	if req.TaskName != "" {
		db = db.Where("st.task_name LIKE ?", "%"+req.TaskName+"%")
	}
	if req.TaskStatus != "" {
		db = db.Where("st.task_status = ?", req.TaskStatus)
	}
	db.Count(&total)
	if req.PageNum > 0 && req.PageSize > 0 {
		offset := cwrs_utils.CalcOffset(req.PageNum, req.PageSize)
		db = db.Offset(offset).Limit(req.PageSize)
	}
	if err := db.Order("st.created_time").Find(&list).Error; err != nil {
		sysTaskLog.Error("GetSysTaskList Error", zap.Error(err))
		return nil, 0, err
	}
	return list, total, nil
}
