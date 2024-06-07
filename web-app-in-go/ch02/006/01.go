package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
	defer cancel()
	go func() { fmt.Println("another go routine") }()
	fmt.Println("STOP")
	<-ctx.Done()
	fmt.Println("and time goes on...")
}
