package main

import (
	"context"
	"github.com/patyukin/go-online-library/internal/config"
	"github.com/patyukin/go-online-library/internal/cronjob"
	"github.com/patyukin/go-online-library/internal/repository"
	"github.com/patyukin/go-online-library/internal/server"
	"github.com/patyukin/go-online-library/internal/service/promotion"
	"github.com/patyukin/go-online-library/pkg/db/mysql"
	"github.com/patyukin/go-online-library/pkg/db/transaction"
	"github.com/patyukin/go-online-library/pkg/migrator"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg, err := config.Get("./internal/config/env.conf")
	if err != nil {
		log.Fatal(err)
	}

	dbClient, err := mysql.New(ctx, cfg.MySQL.DSN)
	err = migrator.UpMigrations(dbClient.GetSqlDB())
	txManager := transaction.NewTransactionManager(dbClient.DB())

	filterRepo := repository.NewFilterRepo(dbClient)
	promotionRepo := repository.NewPromotionRepo(dbClient)
	directoryRepo := repository.NewDirectoryRepo(dbClient)
	promotionService := promotion.New(promotionRepo, filterRepo, directoryRepo, txManager)
	//handlers := server.NewHandler(promotionService)
	handlers := server.NewHandler(promotionService)

	errSrvCh := make(chan error)
	srv := server.New(handlers)
	go func() {
		if err = srv.Run("0.0.0.0:8087"); err != nil {
			logrus.Fatalf("error occured while running http server: %v", err)
			errSrvCh <- err
		}
	}()

	errCronErr := make(chan error)
	cj := cronjob.NewCronJob()
	err = cj.AddFunc("@every 1h", func() {
		err = promotionService.GetAllPromotions(ctx)
		if err != nil {
			logrus.Fatalf("error occured while adding cron job: %v", err)
			errCronErr <- err
		}
	})

	if err != nil {
		logrus.Fatalf("error occured while adding cron job: %v", err)
	}

	go func() {
		cj.Start()
	}()

	logrus.Print("TodoApp Started")

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT)

	select {
	case err = <-errCronErr:
		logrus.Error("Failed to run")
	case err = <-errSrvCh:
		logrus.Error("Failed to run")
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
