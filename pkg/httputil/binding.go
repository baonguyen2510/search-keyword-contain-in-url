package httputil

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func BindJSON[B interface{}](ctx *gin.Context) (*B, error) {
	var body B
	if err := ctx.ShouldBindJSON(&body); err != nil {
		return nil, fmt.Errorf("failed bind json %s", err.Error())
	}
	return &body, nil
}

func BindQuery[B interface{}](ctx *gin.Context) (*B, error) {
	var query B
	if err := ctx.ShouldBindQuery(&query); err != nil {
		return nil, fmt.Errorf("failed bind query %s", err.Error())
	}
	return &query, nil
}
