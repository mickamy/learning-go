package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	s, sep := "", ""

	for i, arg := range os.Args[1:] {
		s = s + sep + strconv.Itoa(i) + ": " + arg
		sep = ", "
	}
	fmt.Println(s)
}
