/* Exercício 1.12: Modifique o servidor Lissajous para ler valores de parâmetros
do URL. Por exemplo, você pode organizá-lo de modo que um URL como
http://localhost:8000/?cycles=20 defina o número de ciclos para 20,
em vez de usar o default igual a 5. Utilize a função strconv.Atoi para
converter o parâmetro do tipo string em um inteiro. Você pode ver a
documentação da função usando go doc strconv.Atoi. */

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
			continue
		}

		lissajousValues[key] = valFloat

	}

	cycles := lissajousValues["cycles"]
	res := lissajousValues["res"]
	size := lissajousValues["size"]
	nframes := lissajousValues["nframes"]
	delay := lissajousValues["delay"]
	intSize := int(size)
	intNframes := int(nframes)
	intDelay := int(delay)

	var palette = []color.Color{color.RGBA{0, 0, 0, 0xff}, color.RGBA{0xff, 0, 0, 0xff}, color.RGBA{0, 0xff, 0, 0xff}, color.RGBA{0, 0, 0xff, 0xff}}

	freq := rand.Float64() * 3.0 // frequência relativa do oscilador y
	anim := gif.GIF{LoopCount: intNframes}
	phase := 0.0 // diferença de fase
	for i := 0; i < intNframes; i++ {
		rect := image.Rect(0, 0, 2*intSize+1, 2*intSize+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(intSize+int(x*size+0.5), intSize+int(y*size+0.5),
				uint8(rand.Intn(len(palette))))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, intDelay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTA: ignorando erros de codificação
}
