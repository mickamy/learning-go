package main

import "fmt"

func main() {
	var i any
	i = 20
	fmt.Println(i)
	i = struct {
		FirstName string
		LastName  string
	}{
		FirstName: "John",
		LastName:  "Lennon",
	}
	fmt.Println(i)
}
