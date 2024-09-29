package db

import (
	"context"
	"log"
	"search-keyword-service/pkg/driver/postgresql"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var defaultStorage *gorm.DB

func MustInit(storage *gorm.DB, err error) {
	if err != nil {
		log.Panic(err.Error())
	}
	defaultStorage = storage
}

// New creates connections to all databases will be used in the application
func New(connection postgresql.Connection, logSQL bool) (*gorm.DB, error) {
	// init mysql connection
	db, err := postgresql.New(connection)
	if err != nil {
		log.Panic(err)
	}
	if logSQL {
		db.Logger = db.Logger.LogMode(logger.Info)
	} else {
		db.Logger = db.Logger.LogMode(logger.Error)
	}

	return db, nil
}

func GetConn() *gorm.DB {
	return defaultStorage
}

func DBWithCtx(ctx context.Context) *gorm.DB {
	if tx, ok := ctx.Value(DbTransactionCtxKey).(*gorm.DB); ok && tx != nil {
		return tx.WithContext(ctx)
	}
	return defaultStorage.WithContext(ctx)
}
