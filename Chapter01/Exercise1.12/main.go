// In the gopl example program, cycles will convent from untyped  to float64 before calculation
// Because const variables in go is untyped,
// we need to do type conversions explicitly when we use a non-consr variable in the program

// Run with "go run main.go web"
// In the browser access "http://localhost:8000/?cycles=20"

package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"
)

var palette = []color.Color{color.White, color.Black}

const (
	greenIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

func main() {

	rand.Seed(time.Now().UTC().UnixNano())
	if len(os.Args) > 1 && os.Args[1] == "web" {

		http.HandleFunc("/", handler)

		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}

}

func handler(w http.ResponseWriter, r *http.Request) {
	u, err := url.Parse(r.URL.RequestURI())
	if err != nil {
		panic(err)
	}
	params := u.Query()
	cycstr := params.Get("cycles")
	cysc, _ := strconv.Atoi(cycstr) //ignore error

	lissajous(w, cysc)
}

func lissajous(out io.Writer, cycs int) {
	//
	cycles := float64(cycs) // number of complete x oscillator revolutions
	const (
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
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
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
