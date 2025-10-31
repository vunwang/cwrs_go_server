package cwrs_zap_logger

import (
	"cwrs_go_server/src/cwrs_core/cwrs_gorm"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"strings"
	"time"
)

// AddSysOperLogReq 新增操作日志入参
type AddSysOperLogReq struct {
	OperId     string `gorm:"column:oper_id" json:"operId"`          // 日志主键
	OperUserId string `gorm:"column:oper_user_id" json:"operUserId"` // 操作用户id
	Method     string `gorm:"column:method" json:"method"`           // 请求方法
	Path       string `gorm:"column:path" json:"path"`               // 请求路径
	Ip         string `gorm:"column:ip" json:"ip"`                   // 客户端IP
	Status     uint   `gorm:"column:status" json:"status"`           // 响应状态码
	ReqBody    string `gorm:"column:req_body" json:"reqBody"`        // 请求体
	ResBody    string `gorm:"column:res_body" json:"resBody"`        // 响应体
	Latency    string `gorm:"column:latency" json:"latency"`         // 耗时
	OperTime   string `gorm:"column:oper_time" json:"operTime"`      // 操作时间
}

func OperationLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 如果是 GET/HEAD/OPTIONS，或没有 Body，跳过
		if c.Request.Method == "GET" || c.Request.Method == "POST" || c.Request.Method == "PUT" || c.Request.Method == "DELETE" {
			if c.FullPath() != "" {
				start := time.Now()

				c.Next()

				latency := int64(time.Since(start) / time.Millisecond)

				userId, _ := c.Get("userId")
				reqBody, _ := c.Get("req_body")
				resBody, _ := c.Get("res_body")
				status, _ := c.Get("status")
				// 生成一个新的UUID v4
				newUUID, err := uuid.NewRandom()
				if err != nil {
					ZapLogger.Error("获取UUID失败:", zap.Error(err))
				}

				// 转换为字符串，并移除连字符
				uuidStr := strings.ReplaceAll(newUUID.String(), "-", "")

				var req AddSysOperLogReq
				req.Method = c.Request.Method
				req.Path = c.FullPath()
				req.Ip = c.ClientIP()
				req.Status = getUint(status)
				if c.Request.Method == "POST" || c.Request.Method == "PUT" {
					req.ReqBody = getString(reqBody)
				} else {
					query := c.Request.URL.Query()
					//Query 参数 ?page=1&size=10
					req.ReqBody = query.Encode()
				}
				req.ResBody = getString(resBody)
				req.Latency = fmt.Sprintf("%dms", latency)
				req.OperId = uuidStr
				req.OperUserId = getString(userId)
				req.OperTime = time.Now().Format("2006-01-02 15:04:05")

				// 异步写入数据库（推荐，避免阻塞请求）
				go func() {
					if req.OperUserId != "" {
						if err := cwrs_gorm.GormDb.Table("sys_oper_log").Create(&req).Error; err != nil {
							ZapLogger.Error("写入操作日志失败", zap.Error(err))
						}
					}
				}()
			}
		}
	}
}

// 辅助函数：安全类型转换
func getUint(v interface{}) uint {
	if u, ok := v.(uint); ok {
		return u
	}
	if u, ok := v.(float64); ok {
		return uint(u)
	}
	if u, ok := v.(int); ok {
		return uint(u)
	}
	return 0
}

func getString(v interface{}) string {
	if s, ok := v.(string); ok {
		return s
	}
	return ""
}
