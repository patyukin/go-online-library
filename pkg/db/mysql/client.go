package mysql

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/patyukin/go-online-library/pkg/db"
)

type Client struct {
	dbClient db.DB
}

func New(ctx context.Context, dsn string) (*Client, error) {
	dbConn, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to db: %v", err)
	}

	err = dbConn.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return &Client{
		dbClient: &mysql{dbConn: dbConn},
	}, nil
}

func (c *Client) DB() db.DB {
	return c.dbClient
}

func (c *Client) Close() error {
	if c.dbClient != nil {
		c.dbClient.Close()
	}

	return nil
}

func (c *Client) GetSqlDB() *sql.DB {
	return c.dbClient.GetSqlDB()
}
