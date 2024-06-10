package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.New(rand.NewSource(time.Now().Unix()))
	a := rand.Intn(10)
	for a < 100 {
		fmt.Println(a)
		if a%5 == 0 {
			goto done
		}

		a = a*2 + 1
	}
	fmt.Println("ループが以上終了した時に行う処理を実行")
done:
	fmt.Println("ループが終わった時に必ず行う処理を実行")
	fmt.Println(a)
}
