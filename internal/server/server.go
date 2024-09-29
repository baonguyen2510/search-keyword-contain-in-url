package server

import (
	"context"
	"errors"
	"search-keyword-service/configs"
	"search-keyword-service/internal/repository/db"
	"search-keyword-service/internal/repository/redis"
	"search-keyword-service/internal/routes"
	"search-keyword-service/pkg/cache"
	"search-keyword-service/pkg/cache/credis"
	"search-keyword-service/pkg/driver/postgresql"
	"search-keyword-service/pkg/http"
	"search-keyword-service/pkg/log"
)

// Instance represents an instance of the server
type Instance struct {
	ctx        context.Context
	name       string
	httpServer *http.HTTPServer
}

// NewInstance returns a new instance of our server
func NewInstance(name string) *Instance {
	return &Instance{name: name}
}

// Start starts the server
func (i *Instance) Start(ctx context.Context) {
	i.ctx = ctx
	var err error
	db.MustInit(db.New(postgresql.Connection{
		ConnectionName: configs.Config.DbConnection,
		Host:           configs.Config.DbHost,
		Port:           configs.Config.DbPort,
		Username:       configs.Config.DbUsername,
		Password:       configs.Config.DbPassword,
		DatabaseName:   configs.Config.DbName,
		Schema:         configs.Config.DbSchema,
	}, configs.Config.DbLogSQL))

	redisClient, err := redis.New(i.ctx)
	if err != nil {
		log.Fatalw("can not init redis", "error", err.Error())
	}
	log.Info("redis initialized successfully")

	cacheRedis := credis.NewClient(redisClient)
	err = cache.Init(cacheRedis)
	if err != nil {
		log.Fatalw("can not init cache", "error", err.Error())
	}
	log.Info("cache initialized successfully")

	router := http.NewHTTPRouter(
		routes.Recover(),
		routes.CORS(),
		routes.Logging(),
		routes.Trace(),
		routes.GetRoutes(),
	)

	// Startup the HTTP HTTPServer in a way that we can gracefully shut it down again
	i.httpServer = http.NewHTTPServer(
		http.HTTPServerWithName(i.name),
		http.HTTPServerWithAddress(
			configs.Config.HttpHost,
			configs.Config.HttpPort,
		),
		http.HTTPServerWithHandler(router))
	err = i.httpServer.Run()
	if !errors.Is(err, http.ErrServerClosed) {
		log.Errorw("HTTP HTTPServer stopped unexpected", "error", err)
		i.Shutdown(ctx)
	}
}

// Shutdown stops the server
func (i *Instance) Shutdown(ctx context.Context) {
	if err := redis.GetClient().Close(); err != nil {
		log.Infof("Failed close redis err %v", err)
	}
	// Shutdown HTTP server
	i.httpServer.Stop(ctx)
}
