package main

import (
	"context"
	"log/slog"
	"net/http"
	"net/http/pprof"
	"os"
	"os/signal"
	"runtime"
	"sync"
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
		mx:     new(sync.Mutex),
		data:   make(map[string]int),
	}

	mux := http.NewServeMux()
	mux.HandleFunc("POST /v1/action", s.endPoint)

	runtime.SetMutexProfileFraction(1)
	mux.Handle("/debug/pprof/mutex", pprof.Handler("mutex"))

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
	mx     *sync.Mutex
	data   map[string]int
}

func (s *server) endPoint(w http.ResponseWriter, r *http.Request) {
	s.logger.Info("request")

	s.writeF1()
	s.writeF2()
	s.readF3()
	s.writeF4()
	s.writeF5()

	w.WriteHeader(http.StatusOK)
}

//go:noinline
func (s *server) writeF1() {
	s.mx.Lock()
	defer s.mx.Unlock()

	time.Sleep(time.Millisecond * 10)
	s.data["writeF1"]++
}

//go:noinline
func (s *server) writeF2() {
	s.mx.Lock()
	defer s.mx.Unlock()

	time.Sleep(time.Millisecond * 10)
	s.data["writeF2"]++
}

//go:noinline
func (s *server) readF3() int {
	s.mx.Lock()
	defer s.mx.Unlock()

	time.Sleep(time.Millisecond * 200)
	return s.data["readF3"]
}

//go:noinline
func (s *server) writeF4() {
	s.mx.Lock()
	defer s.mx.Unlock()

	time.Sleep(time.Millisecond * 10)
	s.data["writeF4"]++
}

//go:noinline
func (s *server) writeF5() {
	s.mx.Lock()
	defer s.mx.Unlock()

	time.Sleep(time.Millisecond * 10)
	s.data["writeF5"]++
}
