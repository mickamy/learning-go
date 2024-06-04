package main

func f(xp *int) {
	*xp = 100
}

func main() {
	x := 10
	f(&x)
	println(x)
}
