package main

import "fmt"

type Hoge struct{ n int }
type Bar struct{ n int }

func (h Hoge) Do() { fmt.Println(h.n) }
func (b Bar) Do()  { fmt.Println(b.n) }

type Fuga struct {
	Hoge
	Bar
	n int
}

func main() {
	f := Fuga{Hoge: Hoge{n: 100}, Bar: Bar{n: 300}, n: 200}
	f.Hoge.Do()
	f.Bar.Do()
}
