//go:build unix

package handler

import (
	"context"
	"os/signal"
	"syscall"
)

func getSignalContext(parentCtx context.Context) (context.Context, context.CancelFunc) {
	return signal.NotifyContext(parentCtx, syscall.SIGALRM)
}
