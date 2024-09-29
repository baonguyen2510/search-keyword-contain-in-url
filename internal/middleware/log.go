package middleware

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"search-keyword-service/common"
	"search-keyword-service/pkg/httputil"
	"search-keyword-service/pkg/log"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Logging() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		request, _ := io.ReadAll(c.Request.Body)
		c.Request.Body = io.NopCloser(bytes.NewReader(request))

		c.Next()

		end := time.Now()
		latency := end.Sub(start)
		end = end.UTC()

		fields := []zapcore.Field{
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.Duration("latency", latency),
			zap.String("time", end.Format(time.RFC3339)),
		}

		userID, ok := c.Get(common.ContextKeyUserID)
		if ok {
			fields = append(fields, zap.Any("user-id", userID))
		}

		userEmail, ok := c.Get(common.ContextKeyEmail)
		if ok {
			fields = append(fields, zap.Any("user-email", userEmail))
		}

		response, ok := c.Get(httputil.ResponseKey)
		if ok {
			fields = append(fields, zap.Any("response", response))
		}

		if c.Request.Method != http.MethodGet {
			var m map[string]interface{}
			json.Unmarshal(request, &m)
			fields = append(fields, zap.Any("request", m))
		}

		if httputil.IsResponseSuccess(c) {
			log.Zap().Info(path, fields...)
		} else {
			log.Zap().Debug(path, fields...)
		}
	}
}
