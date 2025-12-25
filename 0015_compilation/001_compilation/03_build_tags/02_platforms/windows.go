//go:build windows && amd64 && go1.24

// //go:build aix || darwin || (js && wasm) || openbsd || solaris || wasip1
// before 1.17: // +build model_test (and -, \n - or)

package handler

import (
	"context"
	"runtime"
)

func getSignalContext(parentCtx context.Context) (context.Context, context.CancelFunc) {
	_ = runtime.GOOS
	_ = runtime.GOARCH

	return context.WithCancel(parentCtx)
}
