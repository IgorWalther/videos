package main

import (
	"context"
	"log/slog"
	"net/http"
	"net/http/pprof"
	"os"
	"os/signal"
	"runtime"
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

	runtime.SetBlockProfileRate(1)
	mux.Handle("/debug/pprof/block", pprof.Handler("block"))

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

	s.f1()
	s.f2()
	s.f3()
	s.f4()
	s.f5()

	w.WriteHeader(http.StatusOK)
}

//go:noinline
func (s *server) f1() {
	select {
	case <-time.Tick(time.Millisecond * 10):
		return
	}
}

//go:noinline
func (s *server) f2() {
	select {
	case <-time.Tick(time.Millisecond * 10):
		return
	}
}

//go:noinline
func (s *server) f3() {
	select {
	case <-time.Tick(time.Millisecond * 200):
		return
	}
}

//go:noinline
func (s *server) f4() {
	select {
	case <-time.Tick(time.Millisecond * 10):
		return
	}
}

//go:noinline
func (s *server) f5() {
	select {
	case <-time.Tick(time.Millisecond * 10):
		return
	}
}
