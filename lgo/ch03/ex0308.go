package main

import "fmt"

func main() {
	x := make([]int, 0, 5)
	x = append(x, 1, 2, 3, 4)
	fmt.Println("x", x)
	y := x[:2:2]
	fmt.Println("y", y)
	z := x[2:4:4]
	fmt.Println("z", z)

	fmt.Println(cap(x), cap(y), cap(z))

	y = append(y, 30, 40, 50)
	x = append(x, 60)
	z = append(z, 70)

	fmt.Println("x", x)
	fmt.Println("y", y)
	fmt.Println("z", z)
}
