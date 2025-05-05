CURDIR=$(shell pwd)
PPROFDIR=${CURDIR}/pprof

.PHONY: blockprof
blockprof:
	(PPROF_TMPDIR=${PPROFDIR} go tool pprof -http :8082 -seconds 20 http://127.0.0.1:8080/debug/pprof/block)
