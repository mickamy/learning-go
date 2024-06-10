package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func fileChecker(name string) (err error) {
	f, err := os.Open(name)
	if err != nil {
		return fmt.Errorf("open file: %w", err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Printf("file close err: %v", err)
		}
	}()
	return nil
}

func main() {
	err := fileChecker("test.txt")
	if err != nil {
		fmt.Println(err)
		if wrappedErr := errors.Unwrap(err); wrappedErr != nil {
			fmt.Println("in main, wrappedErr:", wrappedErr)
		}
	}
}
