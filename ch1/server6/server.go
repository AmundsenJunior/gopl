// server is a minimal echo server that returns lissajous gifs
package main

import (
	"github.com/amundsenjunior/gopl/ch1/server6/lissajous"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	params := lissajous.SetLissajousParams(r.Form)

	lissajous.Lissajous(w, params)
}
