package routes

import (
	"search-keyword-service/configs"
	"search-keyword-service/internal/handlers/keyword"
	"search-keyword-service/internal/middleware"
	"search-keyword-service/pkg/http"
	"time"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

func routes(r *gin.Engine) {
	v1 := r.Group("v1/search")

	if configs.Config.SwaggerEnable {
		v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	v1.GET("/health", func(c *gin.Context) {
		c.Status(200)
	})

	// API Endpoints
	keywordHandler := keyword.New()
	keywordGroup := v1.Group("/keyword", gin.BasicAuth(gin.Accounts{
		configs.Config.BasicAuthUser: configs.Config.BasicAuthPassword,
	}))
	{
		keywordGroup.GET("/rank/:word", keywordHandler.GetKeywordRank)
		// Update keyword by manual
		keywordGroup.POST("/sync/:word", keywordHandler.SyncKeywordRank)
	}

	// Background Scheduler update all keywords qualify in DB
	go func() {
		for {
			time.Sleep(time.Duration(configs.Config.ConfigTimeSchedule) * time.Second) // Configurable schedule
			keywordHandler.SyncAllKeywordsRank()
		}
	}()
}

func GetRoutes() http.HTTPRouterOption {
	return func(router *http.HttpRouter) {
		routes(router.BaseRoute())
	}
}

func CORS() http.HTTPRouterOption {
	return func(router *http.HttpRouter) {
		router.BaseRoute().Use(func(c *gin.Context) {
			origin := c.GetHeader("origin")
			if origin == "" {
				origin = "*"
			}
			c.Header("Access-Control-Allow-Origin", origin)
			c.Header("Access-Control-Allow-Credentials", "true")
			c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
			c.Header("Access-Control-Allow-Methods", "POST, HEAD, PATCH, OPTIONS, GET, PUT, DELETE")
			if c.Request.Method == "OPTIONS" {
				c.AbortWithStatus(204)
				return
			}
			c.Next()
		})
	}
}

func Recover() http.HTTPRouterOption {
	return func(router *http.HttpRouter) {
		router.BaseRoute().Use(gin.Recovery())
	}
}

func Logging() http.HTTPRouterOption {
	return func(router *http.HttpRouter) {
		router.BaseRoute().Use(middleware.Logging())
	}
}

func Trace() http.HTTPRouterOption {
	return func(router *http.HttpRouter) {
		router.BaseRoute().Use(otelgin.Middleware(configs.Config.ServerName))
	}
}
