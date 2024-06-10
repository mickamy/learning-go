package main

import "fmt"

func main() {
	intSet := map[int]bool{}
	vals := []int{4, 20, 2, 5, 8, 76, 3, 8, 9, 2, 1, 10}
	for _, val := range vals {
		intSet[val] = true
	}
	fmt.Println(len(vals), len(intSet))
	fmt.Println(intSet[5])
	fmt.Println(intSet[500])
	if intSet[100] {
		fmt.Println("100 is not set")
	}
	if intSet[10] {
		fmt.Println("10 is set")
	}
}
