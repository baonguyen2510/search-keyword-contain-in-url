package db

import (
	"context"
	"search-keyword-service/pkg/log"

	"gorm.io/gorm"
)

type CtxKey string

const (
	DbTransactionCtxKey CtxKey = "db-transaction-key"
)

type DBTransactionManager interface {
	EndTx(error) error
	RecoverTx()
	AssignToContext(parentCtx context.Context) context.Context
}

func BeginTx() DBTransactionManager {
	tx := defaultStorage.Begin()
	return dbTrx{
		Tx: tx,
	}
}

type dbTrx struct {
	Tx *gorm.DB
}

// RecoverTx to recover & roll back transaction of this service
func (txm dbTrx) RecoverTx() {
	if p := recover(); p != nil {
		log.Error("found p and rollback tx: ", p)
		txm.Tx.Rollback()
		panic(p)
	}
}

// EndTx to end (commit or rollback) the current transaction of this service
func (txm dbTrx) EndTx(err error) error {
	if err != nil {
		log.Error("found e and rollback: ", err)
		txm.Tx.Rollback()
	} else {
		err = txm.Tx.Commit().Error
		if err != nil {
			log.Error("found e when commit and rollback: ", err)
			txm.Tx.Rollback()
		}
	}
	return err
}

// AssignToContext assign transaction to context
func (txm dbTrx) AssignToContext(parentCtx context.Context) context.Context {
	return context.WithValue(parentCtx, DbTransactionCtxKey, txm.Tx)
}
