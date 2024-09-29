package httputil

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

const (
	ResponseKey = "domain/gin-gonic/gin/responsekey"
)

type baseResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Detail  interface{} `json:"detail,omitempty"`
}

func RespondWarpJSON(c *gin.Context, httpStatusCode int, payload interface{}) {
	if payload == nil {
		payload = "success"
	}
	response := &baseResponse{
		Success: true,
		Data:    payload,
		Detail:  nil,
	}
	c.Set(ResponseKey, response)
	c.JSON(httpStatusCode, response)
}

func RespondWrapError(c *gin.Context, httpStatusCode int, message string, detail interface{}) {
	span := trace.SpanFromContext(c)
	span.SetStatus(codes.Error, message)

	if err, ok := detail.(error); ok {
		span.RecordError(err)
	} else {
		span.RecordError(errors.New(message))
	}

	response := &baseResponse{
		Success: false,
		Message: message,
		Detail:  detail,
	}
	c.Set(ResponseKey, response)
	c.JSON(httpStatusCode, response)
}

// 200 <= status < 300
func IsResponseSuccess(c *gin.Context) bool {
	status := c.Writer.Status()
	return status >= http.StatusOK && status < http.StatusMultipleChoices
}
