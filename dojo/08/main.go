package main

import "fmt"

type Hoge interface {
	M()
	N()
}

type HogeImpl struct{}

func (hoge HogeImpl) M() {
	fmt.Println("M")
}

func (hoge HogeImpl) N() {
	fmt.Println("N")
}

type fuga struct{ Hoge }

func (f fuga) M() {
	fmt.Println("fuga")
	f.Hoge.M()
}

func HiHoge(h Hoge) Hoge {
	return fuga{h}
}

func main() {
	var hoge = HogeImpl{}
	HiHoge(hoge).M()
}
