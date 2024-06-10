package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4}
	num := copy(x[:3], x[1:])
	fmt.Println(x, num)
}
