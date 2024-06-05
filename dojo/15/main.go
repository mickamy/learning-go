package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

var seed = rand.NewSource(time.Now().UTC().UnixMilli())
var rdg = rand.New(seed)

func handler(w http.ResponseWriter, r *http.Request) {
	val := omikuji(rdg)
	_, err := fmt.Fprintf(w, val)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func omikuji(rand *rand.Rand) string {
	i := rand.Intn(6) % 6
	switch i {
	case 0:
		return "大吉"
	case 1:
		return "中吉"
	case 2:
		return "小吉"
	case 3:
		return "小凶"
	case 4:
		return "凶"
	case 5:
		return "大凶"
	default:
		panic(fmt.Errorf("unexpected value %s", i))
	}
}
