package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var withLineNumber = flag.Bool("n", false, "Display line number or not (default false)")

func main() {
	flag.Parse()
	files := flag.Args()
	for _, file := range files {
		scanFileAndOut(file)
	}
}

func scanFileAndOut(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening %s: %v\n", filename, err)
		os.Exit(1)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Error closing file:", err)
			os.Exit(1)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	lineNum := 0
	for scanner.Scan() {
		line := scanner.Text()
		lineNum++
		if *withLineNumber {
			fmt.Printf("%d: %s\n", lineNum, scanner.Text())
		} else {
			fmt.Printf("%s\n", line)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
		os.Exit(1)
	}
}
