package middleware

import (
	"bytes"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/okh8609/gin_blog/global"
	"github.com/okh8609/gin_blog/pkg/logger"
)

type AccessLogWriter struct {
	gin.ResponseWriter
	body_buffer *bytes.Buffer // 會用來把 body 也順便寫入到這個 buffer
}

func (w AccessLogWriter) Write(p []byte) (int, error) {
	if n, err := w.body_buffer.Write(p); err != nil { // 把 body 順便寫進去 buffer
		return n, err
	}
	return w.ResponseWriter.Write(p)
}

func AccessLog(c *gin.Context) {
	bodyWriter := &AccessLogWriter{body_buffer: bytes.NewBufferString(""), ResponseWriter: c.Writer}
	c.Writer = bodyWriter

	beginTime := time.Now().Unix()
	c.Next()
	endTime := time.Now().Unix()

	fields := logger.Field{
		"request":  c.Request.PostForm.Encode(),
		"response": bodyWriter.body_buffer.String(), // 把 buffer 中的資料讀出來
	}
	s := "access log: method: %s, status_code: %d, begin_time: %d, end_time: %d"
	global.MyLogger.WithField(fields).Infof(c, s,
		c.Request.Method,
		bodyWriter.Status(),
		beginTime,
		endTime,
	)
}
