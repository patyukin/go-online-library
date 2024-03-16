package mysql

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/patyukin/go-online-library/pkg/db"
)

type mysqlClient struct {
	dbClient db.DB
}

func New(_ context.Context, dsn string) (db.Client, error) {
	dbConn, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to db: %v", err)
	}

	return &mysqlClient{
		dbClient: &mysql{dbConn: dbConn},
	}, nil
}

func (c *mysqlClient) DB() db.DB {
	return c.dbClient
}

func (c *mysqlClient) Close() error {
	if c.dbClient != nil {
		c.dbClient.Close()
	}

	return nil
}

func (c *mysqlClient) GetSqlDB() *sql.DB {
	return c.dbClient.GetSqlDB()
}
