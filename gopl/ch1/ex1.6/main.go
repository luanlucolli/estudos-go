/*Exercício 1.6: Modifique o programa Lissajous para gerar imagens em váriascores adicionando mais valores a palette para então exibi-las alterando o
terceiro argumento de SetColorIndex de alguma maneira interessante.*/

package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

var palette = []color.Color{color.RGBA{0, 0, 0, 0xff}, color.RGBA{0xff, 0, 0, 0xff}, color.RGBA{0, 0xff, 0, 0xff}, color.RGBA{0, 0, 0xff, 0xff}}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	lissajous(os.Stdout)
}
func lissajous(out io.Writer) {
	const (
		cycles  = 5     // número de revoluções completas do oscilador x
		res     = 0.001 // resolução angular
		size    = 300   // canvas da imagem cobre de [-size..+size]
		nframes = 64    // número de quadros da animação
		delay   = 8     // tempo entre quadros em unidades de 10ms
	)
	freq := rand.Float64() * 3.0 // frequência relativa do oscilador y
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // diferença de fase
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				uint8(rand.Intn(len(palette))))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTA: ignorando erros de codificação
}
