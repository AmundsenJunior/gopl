// server is a minimal echo server that returns lissajous gifs
package main

import (
	"github.com/amundsenjunior/gopl/ch1/server5/lissajous"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

type LissajousParams struct {
	cycles, size, nframes, delay int
	res, phase, freqmod          float64
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	params := setLissajousParams(r.Form)

	lissajous.Lissajous(w, params.cycles, params.size, params.nframes, params.delay, params.res, params.phase, params.freqmod)
}

func setLissajousParams(form url.Values) *LissajousParams {
	lissajous_params := new(LissajousParams)

	if _, ok := form["cycles"]; ok {
		if v, err := strconv.Atoi(form.Get("cycles")); err == nil {
			lissajous_params.cycles = v
		}
	} else {
		lissajous_params.cycles = 5
	}

	if _, ok := form["size"]; ok {
		if v, err := strconv.Atoi(form.Get("size")); err == nil {
			lissajous_params.size = v
		}
	} else {
		lissajous_params.size = 100
	}

	if _, ok := form["nframes"]; ok {
		if v, err := strconv.Atoi(form.Get("nframes")); err == nil {
			lissajous_params.nframes = v
		}
	} else {
		lissajous_params.nframes = 64
	}

	if _, ok := form["delay"]; ok {
		if v, err := strconv.Atoi(form.Get("delay")); err == nil {
			lissajous_params.delay = v
		}
	} else {
		lissajous_params.delay = 8
	}

	if _, ok := form["res"]; ok {
		if v, err := strconv.ParseFloat(form.Get("res"), 64); err == nil {
			lissajous_params.res = v
		}
	} else {
		lissajous_params.res = 0.001
	}

	if _, ok := form["phase"]; ok {
		if v, err := strconv.ParseFloat(form.Get("phase"), 64); err == nil {
			lissajous_params.phase = v
		}
	} else {
		lissajous_params.phase = 0.0
	}

	if _, ok := form["freqmod"]; ok {
		if v, err := strconv.ParseFloat(form.Get("freqmod"), 64); err == nil {
			lissajous_params.freqmod = v
		}
	} else {
		lissajous_params.freqmod = 3.0
	}

	return lissajous_params
}
