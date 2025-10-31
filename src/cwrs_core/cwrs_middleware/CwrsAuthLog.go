package cwrs_middleware

import (
	"bytes"
	"cwrs_go_server/src/cwrs_core/cwrs_zap_logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io"
)

// 鉴权接口日志
func SysAuthLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.GetString("userId")
		account := c.GetString("account")
		role := c.GetString("role")

		if c.Request.Method == "POST" || c.Request.Method == "PUT" {
			data, _ := io.ReadAll(c.Request.Body) //ReadAll读取流后,会清空原来的流. 所以下面2步是将数据流恢复回去
			buf := bytes.NewBuffer(data)
			c.Request.Body = io.NopCloser(buf)
			cwrs_zap_logger.ZapLogger.Info(
				"系统日志",
				zap.String("Method", c.Request.Method),
				zap.String("userId", userId),
				zap.String("account", account),
				zap.String("role", role),
				zap.String("url", c.Request.RequestURI),
				zap.String("参数", string(data)),
			)
		} else {
			cwrs_zap_logger.ZapLogger.Info(
				"系统日志",
				zap.String("Method", c.Request.Method),
				zap.String("userId", userId),
				zap.String("account", account),
				zap.String("role", role),
				zap.String("url", c.Request.RequestURI),
			)
		}
	}
}

// 非鉴权接口日志
func SysNotAuthLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("token") != "" {
			return
		}

		if c.Request.Method == "POST" || c.Request.Method == "PUT" {
			data, _ := io.ReadAll(c.Request.Body) //ReadAll读取流后,会清空原来的流. 所以下面2步是将数据流恢复回去
			buf := bytes.NewBuffer(data)
			c.Request.Body = io.NopCloser(buf)
			cwrs_zap_logger.ZapLogger.Info(
				"系统日志",
				zap.String("Method", c.Request.Method),
				zap.String("url", c.Request.RequestURI),
				zap.String("参数", string(data)),
			)
		} else {
			cwrs_zap_logger.ZapLogger.Info(
				"系统日志",
				zap.String("Method", c.Request.Method),
				zap.String("url", c.Request.RequestURI),
			)
		}
	}
}
