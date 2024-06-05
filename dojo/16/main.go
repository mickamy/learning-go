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
 