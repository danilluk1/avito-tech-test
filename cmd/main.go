package cmd

import (
	"context"
	"database/sql"
	"github.com/danilluk1/avito-tech/config"
	api "github.com/danilluk1/avito-tech/internal/api/router"
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

	ctx := context.Background()

	db, err := sql.Open("postgres", cfg.DbConn)
	if err != nil {
		panic(err)
	}

	router := api.Setup()

	srv := &http.Server{
		Addr:         "localhost:3002",
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
