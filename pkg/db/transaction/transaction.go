package transaction

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/patyukin/go-online-library/pkg/db"
	"github.com/patyukin/go-online-library/pkg/db/mysql"
)

type manager struct {
	db db.Transactor
}

func NewTransactionManager(db db.Transactor) db.TxManager {
	return &manager{
		db: db,
	}
}

func (m *manager) transaction(ctx context.Context, opts sql.TxOptions, fn db.Handler) (err error) {
	tx, ok := ctx.Value(mysql.TxKey).(*sql.Tx)
	if ok {
		return fn(ctx)
	}

	tx, err = m.db.BeginTx(ctx, &opts)
	if err != nil {
		return fmt.Errorf("failed to start transaction: %w", err)
	}

	ctx = context.WithValue(ctx, mysql.TxKey, tx)

	defer func() {
		if errRollback := tx.Rollback(); errRollback != nil {
			fmt.Printf("failed to rollback transaction: %w", errRollback)
		}
	}()

	if err = fn(ctx); err != nil {
		return fmt.Errorf("failed to execute transaction: %w", err)
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (m *manager) ReadCommitted(ctx context.Context, f db.Handler) error {
	txOpts := sql.TxOptions{Isolation: sql.LevelReadCommitted}
	return m.transaction(ctx, txOpts, f)
}
