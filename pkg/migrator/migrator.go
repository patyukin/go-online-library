package migrator

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pressly/goose/v3"
	"github.com/sirupsen/logrus"
)

func UpMigrations(db *sql.DB) error {
	if err := goose.SetDialect("mysql"); err != nil {
		return fmt.Errorf("failed to set dialect: %v", err)
	}

	if err := goose.Up(db, "./migrations"); err != nil {
		return fmt.Errorf("failed to up migrations: %v", err)
	}

	logrus.Info("up migrations")

	return nil
}
