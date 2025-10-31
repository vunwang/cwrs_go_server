package pojo

type SysOperLog struct {
	OperId     string `gorm:"column:oper_id" json:"operId"`          // 日志主键
	OperUserId string `gorm:"column:oper_user_id" json:"operUserId"` // 操作用户id
	Method     string `gorm:"column:method" json:"method"`           // 请求方法
	Path       string `gorm:"column:path" json:"path"`               // 请求路径
	Ip         string `gorm:"column:ip" json:"ip"`                   // 客户端IP
	Status     int    `gorm:"column:status" json:"status"`           // 响应状态码
	ReqBody    string `gorm:"column:req_body" json:"reqBody"`        // 请求参数
	ResBody    string `gorm:"column:res_body" json:"resBody"`        // 响应数据
	Latency    string `gorm:"column:latency" json:"latency"`         // 耗时
	OperTime   string `gorm:"column:oper_time" json:"operTime"`      // 操作时间
}

type SysOperLogResp struct {
	OperId     string `gorm:"column:oper_id" json:"operId"`          // 日志主键
	OperUserId string `gorm:"column:oper_user_id" json:"operUserId"` // 操作用户id
	Method     string `gorm:"column:method" json:"method"`           // 请求方法
	Path       string `gorm:"column:path" json:"path"`               // 请求路径
	Ip         string `gorm:"column:ip" json:"ip"`                   // 客户端IP
	OperName   string `gorm:"column:oper_name" json:"operName"`      // 操作用户名称
	Status     int    `gorm:"column:status" json:"status"`           // 响应状态码
	ReqBody    string `gorm:"column:req_body" json:"reqBody"`        // 请求参数
	ResBody    string `gorm:"column:res_body" json:"resBody"`        // 响应数据
	Latency    string `gorm:"column:latency" json:"latency"`         // 耗时
	OperTime   string `gorm:"column:oper_time" json:"operTime"`      // 操作时间
}

type SysOperLogListResp struct {
	OperId     string `gorm:"column:oper_id" json:"operId"`          // 日志主键
	OperUserId string `gorm:"column:oper_user_id" json:"operUserId"` // 操作用户id
	Method     string `gorm:"column:method" json:"method"`           // 请求方法
	Path       string `gorm:"column:path" json:"path"`               // 请求路径
	Ip         string `gorm:"column:ip" json:"ip"`                   // 客户端IP
	OperName   string `gorm:"column:oper_name" json:"operName"`      // 操作用户名称
	Status     int    `gorm:"column:status" json:"status"`           // 响应状态码
	Latency    string `gorm:"column:latency" json:"latency"`         // 耗时
	OperTime   string `gorm:"column:oper_time" json:"operTime"`      // 操作时间
}

// DelSysOperLogReq 删除操作日志入参
type DelSysOperLogReq struct {
	OperIds string `form:"operIds" json:"operIds" binding:"required"` //id 多个英文逗号分隔
}

// GetSysOperLogDetailReq 查询详情入参
type GetSysOperLogDetailReq struct {
	OperId string `form:"operId" json:"operId" binding:"required"` //id
}

// GetSysOperLogListReq 分页查询入参
type GetSysOperLogListReq struct {
	PageNum   int    `form:"pageNum" json:"pageNum"`     // 页码
	PageSize  int    `form:"pageSize" json:"pageSize"`   // 每页显示条数
	OperName  string `form:"operName" json:"operName"`   // 操作用户名称
	StartTime string `form:"startTime" json:"startTime"` // 操作时间开始
	EndTime   string `form:"endTime" json:"endTime"`     // 操作时间结束
}
