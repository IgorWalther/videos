package main

import (
	"net/http"
	"net/http/pprof"
	"runtime"
	"runtime/debug"
)

func main() {
	debug.SetGCPercent(100)
	debug.SetMemoryLimit(100)

	runtime.SetMutexProfileFraction(1)
	runtime.SetBlockProfileRate(1)
	runtime.SetCPUProfileRate(1)
}

func registerProfilerEndpoints(mux *http.ServeMux) {
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
	mux.Handle("/debug/pprof/block", pprof.Handler("block"))
	mux.Handle("/debug/pprof/heap", pprof.Handler("heap"))
	mux.Handle("/debug/pprof/allocs", pprof.Handler("allocs"))
	mux.Handle("/debug/pprof/goroutine", pprof.Handler("goroutine"))
	mux.Handle("/debug/pprof/threadcreate", pprof.Handler("goroutine"))
}
