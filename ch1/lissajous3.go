// lissajous generates gif animations of random lissajous figures.
// WARNING: if on Linux, may only be able to view in-browser, due to
// https://github.com/golang/go/issues/13746, reference to Eye-of-GNOME bug
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

var green = color.RGBA{0x00, 0xFF, 0x00, 0xFF}
var teal = color.RGBA{0x00, 0xFF, 0xFF, 0xFF}
var blue = color.RGBA{0x00, 0x00, 0xFF, 0xFF}
var palette = []color.Color{color.Black, green, teal, blue}

const (
	greenIdx = 1
	tealIdx  = 2
	blueIdx  = 3
)

func main() {
	filename := "output.gif"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}

	output, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0755)
	check(err)
	defer output.Close()

	lissajous(output)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 10    // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 200   // image canvas covers [-size..+size]
		nframes = 256   // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)

		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), colorIndex(i, t))
		}

		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // note: ignoring encoding errors
}

func colorIndex(frame int, increment float64) uint8 {
	control := math.Floor(float64(frame) / increment)
	for {
		switch {
		case control < 2:
			return greenIdx
		case control < 4:
			return tealIdx
		default:
			return blueIdx
		}
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
