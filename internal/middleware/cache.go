package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"search-keyword-service/common"
	"search-keyword-service/pkg/cache"
	"search-keyword-service/pkg/httputil"
	"time"

	"github.com/gin-gonic/gin"
)

func Caching(ttl time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		type Caching struct {
			Response any `json:"response"`
			Status   int `json:"status"`
		}

		key := fmt.Sprintf(common.RedisKeyCacheResponse, c.Request.Method+":"+c.Request.URL.Path+c.Request.URL.RawQuery)
		ctx := c.Request.Context()
		value, err := cache.Get[Caching](ctx, key)
		if err == nil {
			c.Set(httputil.ResponseKey, value.Response)
			c.AbortWithStatusJSON(value.Status, value.Response)
			return
		}

		if !errors.Is(err, cache.ErrNotFound) {
			httputil.RespondWrapError(c, http.StatusBadRequest, err.Error(), nil)
			c.Abort()
			return
		}

		c.Next()

		if httputil.IsResponseSuccess(c) {
			response, ok := c.Get(httputil.ResponseKey)
			if ok {
				value = Caching{
					Response: response,
					Status:   c.Writer.Status(),
				}
				cache.SetEx(ctx, key, value, ttl)
			}
		}
	}
}
