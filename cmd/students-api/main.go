package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Anandusanthosh123/students-api/internal/config"
	"github.com/Anandusanthosh123/students-api/internal/http/handlers/student"

	"github.com/Anandusanthosh123/students-api/internal/storage/sqlite"
)

func main() {
	// load config
	cfg := config.MustLoad()

	// use in built logger package
	// database setup

	storage, err := sqlite.New(cfg)
	if err != nil {
		log.Fatal(err)
	}

	slog.Info("storage initialized ", slog.String("env", cfg.Env), slog.String("version", "1.0.0"))

	// setup router
	router := http.NewServeMux() // returns server mux basically router

	router.HandleFunc("POST /api/students", student.New(storage)) // dependency injection here
	router.HandleFunc("GET /api/students/{id}", student.GetById(storage))

	// setup http server

	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}
	slog.Info("Server Started ", slog.String("address", cfg.Addr))
	// fmt.Printf("Server started %s", cfg.Addr)

	done := make(chan os.Signal, 1) //buffered channel for os signals like interupts

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM) // if there is a signal from os notify in done channel
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("Failed to start server")
		}

	}()
	<-done

	//logic for graceful  server shutdown ongoing - request won't affected,but new request is not accepted
	slog.Info("Shutting down the server") // structured log

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) //passing a empty context and 5sec timeout
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {

		slog.Error("Failed to shutdown server", slog.String("error", err.Error()))

	}

	slog.Info("server shutdown sucessfully")

}
