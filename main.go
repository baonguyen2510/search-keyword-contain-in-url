package main

import (
	"search-keyword-service/configs"
	"search-keyword-service/internal/server"
	"search-keyword-service/pkg/http"
)

const (
	appName = "search-keyword-service"
)

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
