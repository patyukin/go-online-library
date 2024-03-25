package main

import (
	"context"
	"github.com/patyukin/go-online-library/internal/config"
	"github.com/patyukin/go-online-library/internal/cronjob"
	"github.com/patyukin/go-online-library/internal/handler"
	"github.com/patyukin/go-online-library/internal/repository"
	"github.com/patyukin/go-online-library/internal/sender"
	"github.com/patyukin/go-online-library/internal/server"
	"github.com/patyukin/go-online-library/internal/server/router"
	"github.com/patyukin/go-online-library/internal/usecase"
	"github.com/patyukin/go-online-library/pkg/db/mysql"
	"github.com/patyukin/go-online-library/pkg/db/transaction"
	"github.com/patyukin/go-online-library/pkg/migrator"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg, err := config.Get("./internal/config/env.conf")
	if err != nil {
		logrus.Fatalf("error occured while getting config: %v", err)
	}

	dbClient, err := mysql.New(ctx, cfg.MySQL.DSN)
	if err != nil {
		logrus.Fatalf("error occured while connecting to db: %v", err)
	}

	err = migrator.UpMigrations(dbClient.GetSqlDB())
	if err != nil {
		logrus.Fatalf("error occured while up migrations: %v", err)
	}

	txManager := transaction.NewTransactionManager(dbClient.DB())

	repo := repository.New(dbClient)
	sndr := sender.NewSender()
	uc := usecase.New(repo, txManager, sndr)
	h := handler.New(uc)
	rtr := router.Init(h)

	errCh := make(chan error)

	srv := server.New(rtr)

	cj := cronjob.NewCronJob()
	err = cj.AddFunc("@every 1h", func() {
		err = uc.GetAllPromotions(ctx)
		if err != nil {
			logrus.Errorf("error occured while adding cron job: %v", err)
		}
	})

	if err != nil {
		logrus.Fatalf("error occured while adding cron job: %v", err)
	}

	go func() {
		if err = srv.Run("0.0.0.0:8087"); err != nil {
			logrus.Errorf("error occured while running http server: %v", err)
			errCh <- err
		}
	}()

	go func() {
		cj.Start()
	}()

	logrus.Print("Online Library Started")

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT)

	select {
	case err = <-errCh:
		logrus.Errorf("Failed to run, err: %v", err)
	case res := <-sigChan:
		if res == syscall.SIGINT || res == syscall.SIGTERM {
			logrus.Info("Signal received")
		} else if res == syscall.SIGHUP {
			logrus.Info("Signal received")
		}
	}

	cancel()

	logrus.Print("Shutting Down")

	if err = srv.Shutdown(ctx); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}

	if err = dbClient.Close(); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}

	cj.Stop()
}
