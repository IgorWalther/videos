package main

import (
	"context"
	"io"
	"log/slog"
	"math"
	"net/http"
	"net/http/pprof"
	"os"
	"os/signal"
	"runtime"
	"runtime/debug"
	"sync"
	"syscall"
	"time"

	"github.com/grafana/pyroscope-go"
)

const addr = "0.0.0.0:8080"

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	runtime.SetMutexProfileFraction(1)
	runtime.SetBlockProfileRate(1)
	_, err := pyroscope.Start(pyroscope.Config{
		ApplicationName: "leak.app",
		ServerAddress:   "http://pyroscope:4040",
		Logger:          pyroscope.StandardLogger,
		ProfileTypes: []pyroscope.ProfileType{
			pyroscope.ProfileCPU,
			pyroscope.ProfileAllocObjects,
			pyroscope.ProfileAllocSpace,
			pyroscope.ProfileInuseObjects,
			pyroscope.ProfileInuseSpace,

			pyroscope.ProfileGoroutines,
			pyroscope.ProfileMutexCount,
			pyroscope.ProfileMutexDuration,
			pyroscope.ProfileBlockCount,
			pyroscope.ProfileBlockDuration,
		},
	})

	if err != nil {
		panic(err)
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	s := &server{
		logger: logger,
		mx:     new(sync.Mutex),
	}

	mux := http.NewServeMux()
	mux.HandleFunc("POST /v1/action", s.endPoint)

	runtime.SetBlockProfileRate(1)
	mux.Handle("/debug/pprof/block", pprof.Handler("block"))

	server := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	debug.SetGCPercent(-1)
	debug.SetMemoryLimit(math.MaxInt64)

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

	err = server.Shutdown(shutdownCtx)

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
}

func (s *server) endPoint(w http.ResponseWriter, r *http.Request) {
	s.logger.Info("request")

	go s.leak()
}

var global = make([][]byte, 0)

func (s *server) leak() {
	h, _ := http.Get("https://www.google.com")
	body, _ := io.ReadAll(h.Body)

	s.logger.Info("Added", slog.Int("bytes", len(body)))

	s.mx.Lock()
	defer s.mx.Unlock()
	global = append(global, body)

	go s.leak()
}
