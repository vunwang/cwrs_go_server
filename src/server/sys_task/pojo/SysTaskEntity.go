package pojo

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type SysTask struct {
	TaskId        string  `gorm:"column:task_id" json:"taskId"`                // 任务ID
	TaskName      string  `gorm:"column:task_name" json:"taskName"`            // 任务名称
	CronName      string  `gorm:"column:cron_name" json:"cronName"`            // 执行时间(字典-sys_task_cron_expr)
	CronExpr      string  `gorm:"column:cron_expr" json:"cronExpr"`            // Cron表达式-字典值
	FuncName      string  `gorm:"column:func_name" json:"funcName"`            // 要执行的函数名
	TaskParams    JSONMap `gorm:"column:task_params" json:"taskParams"`        // JSON格式参数
	TaskStatus    string  `gorm:"column:task_status" json:"taskStatus"`        // 任务状态(字典-sys_task_status)
	Desc          string  `gorm:"column:desc" json:"desc"`                     // 描述
	LastRunTime   string  `gorm:"column:last_run_time" json:"lastRunTime"`     // 最近一次开启/关闭时间
	NextRunTime   string  `gorm:"column:next_run_time" json:"nextRunTime"`     // 预留字段
	CreatedUserId string  `gorm:"column:created_user_id" json:"createdUserId"` // 创建用户
	CreatedTime   string  `gorm:"column:created_time" json:"createdTime"`      // 创建时间
	UpdatedUserId string  `gorm:"column:updated_user_id" json:"updatedUserId"` // 更新用户
	UpdatedTime   string  `gorm:"column:updated_time" json:"updatedTime"`      // 更新时间
}

type SysTaskResp struct {
	TaskId          string  `gorm:"column:task_id" json:"taskId"`                    // 任务ID
	TaskName        string  `gorm:"column:task_name" json:"taskName"`                // 任务名称
	CronName        string  `gorm:"column:cron_name" json:"cronName"`                // 执行时间(字典-sys_task_cron_expr)
	CronExpr        string  `gorm:"column:cron_expr" json:"cronExpr"`                // Cron表达式-字典值
	FuncName        string  `gorm:"column:func_name" json:"funcName"`                // 要执行的函数名
	TaskParams      JSONMap `gorm:"column:task_params" json:"taskParams"`            // JSON格式参数
	TaskStatus      string  `gorm:"column:task_status" json:"taskStatus"`            // 任务状态(字典-sys_task_status)
	Desc            string  `gorm:"column:desc" json:"desc"`                         // 描述
	LastRunTime     string  `gorm:"column:last_run_time" json:"lastRunTime"`         // 最近一次开启/关闭时间
	CreatedUserId   string  `gorm:"column:created_user_id" json:"createdUserId"`     // 创建用户
	CreatedTime     string  `gorm:"column:created_time" json:"createdTime"`          // 创建时间
	CreatedUserName string  `gorm:"column:created_user_name" json:"createdUserName"` // 创建用户名称
}

// AddSysTaskReq 新增定时任务入参
type AddSysTaskReq struct {
	TaskName   string  `form:"taskName" json:"taskName" binding:"required"` // 任务名称
	FuncName   string  `form:"funcName" json:"funcName" binding:"required"` // 要执行的函数名
	CronName   string  `form:"cronName" json:"cronName" binding:"required"` // 执行时间(字典-sys_task_cron_expr)
	TaskParams JSONMap `form:"taskParams" json:"taskParams"`                // JSON格式参数
	Desc       string  `form:"desc" json:"desc"`                            // 描述
}

// EditSysTaskReq 编辑定时任务入参
type EditSysTaskReq struct {
	TaskId     string  `form:"taskId" json:"taskId" binding:"required"`     // 任务ID
	TaskName   string  `form:"taskName" json:"taskName" binding:"required"` // 任务名称
	FuncName   string  `form:"funcName" json:"funcName" binding:"required"` // 要执行的函数名
	CronName   string  `form:"cronName" json:"cronName" binding:"required"` // 执行时间(字典-sys_task_cron_expr)
	TaskParams JSONMap `form:"taskParams" json:"taskParams"`                // JSON格式参数
	Desc       string  `form:"desc" json:"desc"`                            // 描述
}

// 打开/停止任务入参
type StartAndStopTaskReq struct {
	TaskIds string `form:"taskIds" json:"taskIds" binding:"required"` // id 多个英文逗号分隔
}

// DelSysTaskReq 删除定时任务入参
type DelSysTaskReq struct {
	TaskIds string `form:"taskIds" json:"taskIds" binding:"required"` //id 多个英文逗号分隔
}

// GetSysTaskDetailReq 查询详情入参
type GetSysTaskDetailReq struct {
	TaskId string `form:"taskId" json:"taskId" binding:"required"` //id
}

// GetSysTaskListReq 分页查询入参
type GetSysTaskListReq struct {
	PageNum    int    `form:"pageNum" json:"pageNum"`       // 页码
	PageSize   int    `form:"pageSize" json:"pageSize"`     // 每页显示条数
	TaskName   string `form:"taskName" json:"taskName"`     // 任务名称
	TaskStatus string `form:"taskStatus" json:"taskStatus"` // 任务状态(字典-sys_task_status)
}

type JSONMap map[string]interface{}

// Scan 实现 sql.Scanner
func (j *JSONMap) Scan(src interface{}) error {
	if src == nil {
		*j = nil
		return nil
	}
	var data []byte
	switch v := src.(type) {
	case []byte:
		data = v
	case string:
		data = []byte(v)
	default:
		return fmt.Errorf("不支持的类型: %T", src)
	}
	return json.Unmarshal(data, j)
}

// Value 实现 driver.Valuer
func (j JSONMap) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}
	return json.Marshal(j)
}

// 新增：实现 GORM 的 GormDataTypeInterface
func (JSONMap) GormDataType() string {
	return "json" // 告诉 GORM 用 JSON 类型存
}
