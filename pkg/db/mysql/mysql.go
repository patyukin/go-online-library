package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/georgysavva/scany/v2/sqlscan"
	"github.com/patyukin/go-online-library/pkg/db"
	"github.com/patyukin/go-online-library/pkg/db/prettier"
	"log"
)

type key string

const (
	TxKey key = "tx"
)

type mysql struct {
	dbConn *sql.DB
}

func (p *mysql) ScanOneContext(ctx context.Context, dest interface{}, q db.Query, args ...interface{}) error {
	logQuery(ctx, q, args...)

	row, err := p.QueryContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return sqlscan.ScanOne(dest, row)
}

func (p *mysql) ScanAllContext(ctx context.Context, dest interface{}, q db.Query, args ...interface{}) error {
	logQuery(ctx, q, args...)

	rows, err := p.QueryContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return sqlscan.ScanAll(dest, rows)
}

func (p *mysql) ExecContext(ctx context.Context, q db.Query, args ...interface{}) (sql.Result, error) {
	logQuery(ctx, q, args...)

	tx, ok := ctx.Value(TxKey).(sql.Tx)
	if ok {
		return tx.ExecContext(ctx, q.QueryRaw, args...)
	}

	return p.dbConn.ExecContext(ctx, q.QueryRaw, args...)
}

func (p *mysql) QueryContext(ctx context.Context, q db.Query, args ...interface{}) (*sql.Rows, error) {
	logQuery(ctx, q, args...)

	tx, ok := ctx.Value(TxKey).(sql.Tx)
	if ok {
		return tx.QueryContext(ctx, q.QueryRaw, args...)
	}

	return p.dbConn.QueryContext(ctx, q.QueryRaw, args...)
}

func (p *mysql) QueryRowContext(ctx context.Context, q db.Query, args ...interface{}) *sql.Row {
	logQuery(ctx, q, args...)

	tx, ok := ctx.Value(TxKey).(sql.Tx)
	if ok {
		return tx.QueryRowContext(ctx, q.QueryRaw, args...)
	}

	return p.dbConn.QueryRowContext(ctx, q.QueryRaw, args...)
}

func (p *mysql) BeginTx(ctx context.Context, txOptions *sql.TxOptions) (*sql.Tx, error) {
	return p.dbConn.BeginTx(ctx, txOptions)
}

func (p *mysql) PingContext(ctx context.Context) error {
	return p.dbConn.PingContext(ctx)
}

func (p *mysql) Close() {
	err := p.dbConn.Close()
	if err != nil {
		log.Printf("error occured on db connection close: %v", err)
	}
}

func (p *mysql) GetSqlDB() *sql.DB {
	return p.dbConn
}

func logQuery(ctx context.Context, q db.Query, args ...interface{}) {
	prettyQuery := prettier.Pretty(q.QueryRaw, prettier.PlaceholderDollar, args...)
	log.Println(
		ctx,
		fmt.Sprintf("sql: %s", q.Name),
		fmt.Sprintf("query: %s", prettyQuery),
		fmt.Sprintf("args: %v", args...),
	)
}
