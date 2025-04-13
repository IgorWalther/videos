package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()

	ctxNoCancel := context.WithoutCancel(ctx)
	stub(ctxNoCancel)

	ctxCancel, cancel := context.WithCancel(ctx)
	cancel()
	cancel() // idempotent

	dCtx, c1 := context.WithDeadline(ctx, time.Now().Add(time.Second*27))
	stub(dCtx)
	c1()

	tCtx, c2 := context.WithTimeout(ctx, time.Second*27)
	stub(tCtx)
	c2()

	ctxFunc := context.AfterFunc(ctx, func() {
		fmt.Println("cancelled")
	})
	ctxFunc()

	otherCtx := context.TODO()
	stub(otherCtx)

	ctxValue := context.WithValue(ctxCancel, "id", 27)
	stub(ctxValue)
}

func stub(ctx context.Context) {}
