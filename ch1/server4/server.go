// server is a minimal echo server that returns lissajous gifs
package main

import (
	"log"
	"net/http"
	"github.com/amundsenjunior/gopl/ch1/server4/lissajous"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		lissajous.Lissajous(w)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
