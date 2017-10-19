// lissajous generates gif animations of random lissajous figures.
package lissajous

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
)

var palette = []color.Color{color.White, color.Black}

const blackIndex = 1 // index value pointing to second in palette

//cycles - number of complete x oscillator revolutions
//res - angular resolution
//size - image canvas covers [-size..+size]
//nframes - number of animation frames
//delay - time between frames in 10ms units
//freq_mod - multiplier on frequency
//phase - phase difference
func Lissajous(out io.Writer, cycles, size, nframes, delay int, res, phase, freqmod float64) {
	freq := rand.Float64() * freqmod // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)

		for t := 0.0; t < math.Pi*float64(cycles)*2; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5), blackIndex)
		}

		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // note: ignoring encoding errors
}
