package main

import (
	"context"
	"log/slog"
	"math/rand/v2"
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

func (s *server) endPoint(w http.ResponseWriter, r *http.Request) {
	s.logger.Info("request")

	f1(1)
	f2(2)
	f3(3)
	f4(4)
	f5(5)

	w.WriteHeader(http.StatusOK)
}

func f1(a int) int {
	b := a
	c := a
	d := rand.IntN(2)
	v := c + d
	n := 1 + a + d + rand.IntN(1)

	for range int(10e7) + n + v + b + a {
		a++
	}

	return a
}

func f2(a int) int {
	b := a
	c := a
	d := rand.IntN(2)
	v := c + d
	n := 1 + a + d + rand.IntN(1)

	for range int(10e7) + n + v + b + a {
		a++
	}

	return a
}

func f3(a int) int {
	b := a
	c := a
	d := rand.IntN(2)
	v := c + d
	n := 1 + a + d + rand.IntN(1)

	for range int(10e7) + n + v + b + a {
		a++
	}

	return a
}

func f4(a int) int {
	b := a
	c := a
	d := rand.IntN(2)
	v := c + d
	n := 1 + a + d + rand.IntN(1)

	for range int(10e7) + n + v + b + a {
		a++
	}

	return a
}

func f5(a int) int {
	b := a
	c := a
	d := rand.IntN(2)
	v := c + d
	n := 1 + a + d + rand.IntN(1)

	for range int(10e7) + n + v + b + a {
		a++
	}

	return a
}
