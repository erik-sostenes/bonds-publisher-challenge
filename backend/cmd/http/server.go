package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/erik-sostenes/bonds-publisher-challenge/cmd/bootstrap"
	"github.com/erik-sostenes/bonds-publisher-challenge/cmd/bootstrap/db"
	"github.com/erik-sostenes/bonds-publisher-challenge/cmd/http/health"
)

const defaultPort = "8080"

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	slog.SetDefault(logger)

	port := strings.TrimSpace(os.Getenv("PORT"))
	if port == "" {
		port = defaultPort
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/status", health.HealthCheck())

	dbPtgSQL := db.PostgreSQLInjector()
	bootstrap.BondInjector(dbPtgSQL, mux)
	bootstrap.UserInjector(dbPtgSQL, mux)

	svr := http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	go func() {
		slog.Info("server listening", "port", port)
		log.Fatal(svr.ListenAndServe())
	}()

	signalCH := make(chan os.Signal, 1)
	signal.Notify(signalCH, os.Interrupt)
	<-signalCH

	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()

	_ = svr.Shutdown(ctx)
}
