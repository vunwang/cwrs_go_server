package cwrs_zap_logger

import (
	"bytes"
	"io"
	"strings"

	"github.com/gin-gonic/gin"
)

type bodyReader struct {
	io.ReadCloser
	body []byte
}

func (br *bodyReader) Read(p []byte) (n int, err error) {
	if len(br.body) > 0 {
		n = copy(p, br.body)
		br.body = br.body[n:]
		if len(br.body) == 0 {
			err = io.EOF
		}
		return n, err
	}
	return br.ReadCloser.Read(p)
}

func BodyReader() gin.HandlerFunc {
	return func(c *gin.Context) {
		// GET/HEAD 请求通常无 Body，跳过
		if c.Request.Method == "GET" || c.Request.Method == "HEAD" {
			c.Next()
			return
		}

		// 只读取 JSON / Text 类型（避免 multipart/form-data 导致错误）
		contentType := c.Request.Header.Get("Content-Type")
		if !strings.Contains(contentType, "application/json") &&
			!strings.Contains(contentType, "text/plain") {
			c.Next()
			return
		}

		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.Next()
			return
		}
		c.Request.Body.Close()

		// 重新包装，使其可再次读取
		c.Request.Body = &bodyReader{
			ReadCloser: io.NopCloser(bytes.NewReader(body)),
			body:       body,
		}

		// 存入 Context，供后续使用
		c.Set("req_body", string(body))

		c.Next()
	}
}
