package migrator

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pressly/goose/v3"
)

func UpMigrations(db *sql.DB) error {
	if err := goose.SetDialect("mysql"); err != nil {
		return err
	}

	if err := goose.Up(db, "./migrations"); err != nil {
		return err
	}

	return nil
}
