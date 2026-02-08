package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGALRM)
	defer stop()

	p, _ := os.FindProcess(os.Getpid())
	_ = p.Signal(syscall.SIGALRM)

	<-ctx.Done()
	fmt.Println("err =", ctx.Err())
	fmt.Println("cause =", context.Cause(ctx)) // cause = interrupt signal received
}
