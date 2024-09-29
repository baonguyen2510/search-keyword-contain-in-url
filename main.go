package main

import (
	"search-keyword-service/configs"
	"search-keyword-service/internal/server"
	"search-keyword-service/pkg/http"
)

const (
	appName = "search-keyword-service"
)

// @title Search Keyword Service
// @version 1.0
// @description Service for search, update ranking of keyword.
// @host localhost:7003
// @BasePath /v1/search
// @schemes		http
func main() {
	configs.Init()

	app := http.NewApp(
		http.AppWithName(appName),
		http.AppWithLogger(
			string(configs.Config.AppEnv),
			configs.Config.LoggerDebug,
			configs.Config.LoggerSensitive,
		),
		http.AppWithAction(
			server.NewInstance(appName),
		),
	)

	app.Run()
}
