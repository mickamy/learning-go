package main

import "fmt"

func doTypeSwitch(i any) {
	switch i := i.(type) {
	case nil:
		fmt.Println("nil ", i)
	case int:
		fmt.Println("int ", i)
	case string:
		fmt.Println("string ", i)
	case bool:
		fmt.Println("bool ", i)
	case rune:
		fmt.Println("rune ", i)
	default:
		fmt.Println("unknown ", i)
	}
}

func main() {
	doTypeSwitch(1)
	doTypeSwitch(1.0)
	doTypeSwitch("string")
	doTypeSwitch(true)
	type MyInt int
	var myInt MyInt = 1
	doTypeSwitch(myInt)
}
