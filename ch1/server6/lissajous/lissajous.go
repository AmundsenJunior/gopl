// lissajous generates gif animations of random lissajous figures.
package lissajous

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"net/url"
	"strconv"
)

/*
cycles - number of complete x oscillator revolutions
res - angular resolution
size - image canvas covers [-size..+size]
nframes - number of animation frames
delay - time between frames in 10ms units
freq_mod - multiplier on frequency
phase - phase difference
 */
type LissajousParams struct {
	cycles, size, nframes, delay int
	res, phase, freqmod          float64
}

var palette = []color.Color{color.Black, color.White}

const whiteIndex = 1 // index value pointing to second in palette

func Lissajous(out io.Writer, ljp *LissajousParams) {
	freq := rand.Float64() * ljp.freqmod // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: ljp.nframes}

	for i := 0; i < ljp.nframes; i++ {
		rect := image.Rect(0, 0, 2*ljp.size+1, 2*ljp.size+1)
		img := image.NewPaletted(rect, palette)

		for t := 0.0; t < math.Pi*float64(ljp.cycles)*2; t += ljp.res {
			x := math.Sin(t)
			y := math.Sin(t*freq + ljp.phase)
			img.SetColorIndex(ljp.size+int(x*float64(ljp.size)+0.5), ljp.size+int(y*float64(ljp.size)+0.5), whiteIndex)
		}

		ljp.phase += 0.1
		anim.Delay = append(anim.Delay, ljp.delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // note: ignoring encoding errors
}

func SetLissajousParams(form url.Values) *LissajousParams {
	lissajousParams := new(LissajousParams)

	lissajousParams.cycles = setLissajousIntValue("cycles", 5, form)
	lissajousParams.size = setLissajousIntValue("size", 100, form)
	lissajousParams.nframes = setLissajousIntValue("nframes", 64, form)
	lissajousParams.delay = setLissajousIntValue("delay", 8, form)
	lissajousParams.res = setLissajousFloat64Value("res", 0.001, form)
	lissajousParams.phase = setLissajousFloat64Value("phase", 0.0, form)
	lissajousParams.freqmod = setLissajousFloat64Value("freqmod", 3.0, form)

	return lissajousParams
}

func setLissajousIntValue(keyName string, defaultValue int, form url.Values) int {
	value := defaultValue
	if _, ok := form[keyName]; ok {
		if v, err := strconv.Atoi(form.Get(keyName)); err == nil {
			value = v
		}
	}
	return value
}

func setLissajousFloat64Value(keyName string, defaultValue float64, form url.Values) float64 {
	value := defaultValue
	if _, ok := form[keyName]; ok {
		if v, err := strconv.ParseFloat(form.Get(keyName), 64); err == nil {
			value = v
		}
	}
	return value
}
