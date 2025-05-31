package database

import (
	"context"
	"database/sql"
	"github.com/uptrace/bun"
)

type pgTxManager struct {
	db *bun.DB
}

func NewPostgresTxManager(db *bun.DB) TxManager {
	return &pgTxManager{
		db: db,
	}
}

func (tm *pgTxManager) Begin(ctx context.Context) (bun.Tx, error) {
	return tm.db.BeginTx(ctx, &sql.TxOptions{})
}

func (tm *pgTxManager) Commit(tx bun.Tx) error {
	return tx.Commit()
}

func (tm *pgTxManager) Rollback(tx bun.Tx) error {
	return tx.Rollback()
}
