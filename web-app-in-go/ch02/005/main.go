package main

import (
	"context"
	"fmt"
)

func child(ctx context.Context) {
	if err := ctx.Err(); err != nil {
		fmt.Println("context error:", err)
		return
	}
	fmt.Println("not canceled")
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	child(ctx)
	cancel()
	child(ctx)
}
