package main

import (
	"context"
	"github.com/danilluk1/avito-tech/config"
	"github.com/danilluk1/avito-tech/internal/app/api"
	router "github.com/danilluk1/avito-tech/internal/app/api/router"
	announcementimpl "github.com/danilluk1/avito-tech/internal/services/announcements/impl"
	loggerimpl "github.com/danilluk1/avito-tech/internal/services/logger/impl"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	cfg, err := config.New(true)
	if err != nil {
		panic(err)
	}
	dbConnOpts, err := pq.ParseURL(cfg.DbConn)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	db, err := sqlx.Connect("postgres", dbConnOpts)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	app := &api.App{
		AnnouncementService: announcementimpl.NewAnnouncementService(db),
		Logger:              loggerimpl.NewLogger(),
	}

	router := router.Setup(app)

	srv := &http.Server{
		Addr:         "0.0.0.0:3002",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c

	ctx, cancel := context.WithTimeout(ctx, time.Second*15)
	defer cancel()

	srv.Shutdown(ctx)
	log.Println("shutting down")
	os.Exit(0)

}
