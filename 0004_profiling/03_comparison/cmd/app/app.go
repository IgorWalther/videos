package main

import (
	"context"
	"log/slog"
	"net/http"
	"net/http/pprof"
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
	mux.HandleFunc("POST /v1/action", s.endPoint)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)

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

const iterations = 10_000_000

func (s *server) endPoint(w http.ResponseWriter, r *http.Request) {
	s.logger.Info("request")

	f1()
	f2()
	f3()
	f4()
	f5()

	w.WriteHeader(http.StatusOK)
}

//go:noinline
func f1() int {
	a := 0
	for range iterations {
		a++
	}

	return a
}

//go:noinline
func f2() int {
	a := 0
	for range iterations {
		a++
	}

	return a
}

//go:noinline
func f3() int {
	a := 0

	for range iterations {
		a++
	}

	return a
}

//go:noinline
func f4() int {
	a := 0

	for range iterations {
		a++
	}

	return a
}

//go:noinline
func f5() int {
	a := 0

	for range iterations {
		a++
	}

	return a
}
