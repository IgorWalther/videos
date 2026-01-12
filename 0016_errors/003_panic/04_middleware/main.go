package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const addr = "0.0.0.0:8080"

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	s := &server{
		logger: logger,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("POST /v1/tghook", s.telegramWebhook)
	mux.HandleFunc("POST /v1/panic", s.panic)

	server := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	go func() {
		defer cancel()
		if err := server.ListenAndServe(); err != nil {
			logger.LogAttrs(
				ctx,
				slog.LevelError,
				"server serve error",
				slog.Any("error", err),
			)
		}
	}()

	<-ctx.Done()
	shutdownCtx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	err := server.Shutdown(shutdownCtx)

	if err != nil {
		logger.LogAttrs(
			ctx,
			slog.LevelError,
			"server shutdown error",
			slog.Any("error", err),
		)
		return
	}
}

type server struct {
	logger *slog.Logger
}

func (s *server) telegramWebhook(w http.ResponseWriter, r *http.Request) {
	s.logger.Info("request")

	w.Write([]byte("OK\n"))
	w.WriteHeader(http.StatusOK)
}

// http: panic serving

func (s *server) panic(w http.ResponseWriter, r *http.Request) {
	s.logger.Info("request")

	panic("123")
}
