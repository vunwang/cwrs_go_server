package cwrs_zap_logger

import (
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

type responseInterceptor struct {
	gin.ResponseWriter
	recorder *httptest.ResponseRecorder
}

func (r *responseInterceptor) Write(b []byte) (int, error) {
	r.recorder.Write(b)
	return r.ResponseWriter.Write(b)
}

func (r *responseInterceptor) WriteHeader(code int) {
	r.recorder.WriteHeader(code)
	r.ResponseWriter.WriteHeader(code)
}

func ResponseInterceptor() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 创建 recorder
		recorder := httptest.NewRecorder()

		// 替换 Writer
		ri := &responseInterceptor{
			ResponseWriter: c.Writer,
			recorder:       recorder,
		}
		c.Writer = ri

		c.Next()

		// 获取响应体
		resBody := recorder.Body.String()
		status := recorder.Code

		c.Set("status", status)
		c.Set("res_body", resBody)
	}
}
