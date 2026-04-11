// Server2 é um servidor mínimo de "eco" e contador
package main

import (
	//"fmt"
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		lissajous(w, r)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func lissajous(out io.Writer, r *http.Request) {

	if r.URL.Path == "/favicon.ico" {
		return
	}

	//vou usar um map pra comparar com os valores que vierem no query params
	lissajousValues := map[string]float64{
		"cycles":  5,
		"res":     0.001,
		"size":    300,
		"nframes": 64,
		"delay":   8,
	}

	query := r.URL.Query()

	for key := range lissajousValues {

		valStr := query.Get(key)

		if valStr == "" {
			continue
		}

		valFloat, err := strconv.ParseFloat(valStr, 64)
		if err != nil {
			log.Printf("Erro ao converter query key '%s': %v", key, err)
			break
		}

		lissajousValues[key] = valFloat

		fmt.Println(key, valFloat)
	}

	fmt.Println("")

	intSize := int(lissajousValues["size"])

	var palette = []color.Color{color.RGBA{0, 0, 0, 0xff}, color.RGBA{0xff, 0, 0, 0xff}, color.RGBA{0, 0xff, 0, 0xff}, color.RGBA{0, 0, 0xff, 0xff}}

	freq := rand.Float64() * 3.0 // frequência relativa do oscilador y
	anim := gif.GIF{LoopCount: int(lissajousValues["nframes"])}
	phase := 0.0 // diferença de fase
	for i := 0; i < int(lissajousValues["nframes"]); i++ {
		rect := image.Rect(0, 0, 2*intSize+1, 2*intSize+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < lissajousValues["cycles"]*2*math.Pi; t += lissajousValues["res"] {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(intSize+int(x*lissajousValues["size"]+0.5), intSize+int(y*lissajousValues["size"]+0.5),
				uint8(rand.Intn(len(palette))))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, int(lissajousValues["delay"]))
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTA: ignorando erros de codificação
}
