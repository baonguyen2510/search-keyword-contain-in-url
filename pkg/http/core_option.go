package http

import (
	"context"
	"search-keyword-service/pkg/log"
	"time"
)

type AppOption func(a *App)

// AppWithName set app's name
func AppWithName(name string) AppOption {
	return func(a *App) {
		a.name = name
	}
}

// AppWithContext set contect for app
func AppWithContext(ctx context.Context) AppOption {
	return func(a *App) {
		a.baseCtx = ctx
	}
}

// AppWithLogger init log package instance
func AppWithLogger(mode string, debug bool, sensitiveData bool) AppOption {
	return func(a *App) {
		log.Init(
			log.New(
				log.WithModeFromString(mode),
				log.WithDebug(debug),
				log.EnableSensitive(sensitiveData),
			),
		)
	}
}

// AppWithAction set execute action (start/stop) for process's lifecycle
func AppWithAction(actions AppAction) AppOption {
	return func(a *App) {
		a.startFunc = actions.Start
		a.shutdownFunc = actions.Shutdown
	}
}

// AppWithGracefulTimeout setup graceful timeout
func AppWithGracefulTimeout(waitTime time.Duration) AppOption {
	return func(a *App) {
		a.gracefulTimeout = waitTime
	}
}
