package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go func() {
		time.Sleep(2 * time.Second)
		ch <- 1000
	}()

	select {
	case <-time.After(3 * time.Second):
		fmt.Println("time out")
	case val := <-ch:
		fmt.Println(val)
	}
}
