package database

import (
	"context"
	"github.com/uptrace/bun"
)

type TxManager interface {
	Begin(ctx context.Context) (bun.Tx, error)
	Commit(tx bun.Tx) error
	Rollback(tx bun.Tx) error
}
