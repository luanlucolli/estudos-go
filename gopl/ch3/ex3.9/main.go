/*Exercício 3.9: Escreva um servidor web que renderize fractais e escreva os
dados da imagem ao cliente. Permita que o cliente especifique os valores
de x, y e de zoom como parâmetros da requisição HTTP.*/

package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"net/http"
	"strconv"
)

type FractalRenderContext struct {
	x, y, zoom float64
}

func main() {

	// cada requisição para /mandelbrot chama mandelbrot
	http.HandleFunc("/mandelbrot", renderMandelbrot)
	// inicia o servidor
	log.Fatal(http.ListenAndServe("localhost:8000", nil))

}

func renderMandelbrot(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path == "/favicon.ico" {
		return
	}

	const (
		width  = 600
		height = 320
	)

	// chama a função que lê os dados da URL (isso você já tem)
	frctx := NewFractalRenderContextFromRequest(r)

	larguraJanela := 4.0 / frctx.zoom
	alturaJanela := larguraJanela * float64(height) / float64(width)

	metadeLargura := larguraJanela / 2
	metadeAltura := alturaJanela / 2

	// define dinamicamente os limites aplicando a distância ao redor do centro (x, y)
	xmin := frctx.x - metadeLargura
	xmax := frctx.x + metadeLargura
	ymin := frctx.y - metadeAltura
	ymax := frctx.y + metadeAltura
	w.Header().Set("Content-Type", "image/png")

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Ponto (px, py) da imagem representa o valor complexo z
			img.Set(px, py, mandelbrot(z))

		}
	}
	png.Encode(w, img) // NOTA: ignorando erros

}

// NewFractalRenderContextFromRequest extrai a query, trata os valores e retorna a struct
func NewFractalRenderContextFromRequest(r *http.Request) *FractalRenderContext {

	ctx := &FractalRenderContext{
		x:    -0.5, // Centraliza no meio do fractal
		y:    0.0,
		zoom: 1.0,
	}

	query := r.URL.Query()

	// x, y e zoom
	if x := query.Get("x"); x != "" {
		if val, err := strconv.ParseFloat(x, 64); err == nil {
			ctx.x = val
		}
	}

	if y := query.Get("y"); y != "" {
		if val, err := strconv.ParseFloat(y, 64); err == nil {
			ctx.y = val

		}
	}

	if zoom := query.Get("zoom"); zoom != "" {
		if val, err := strconv.ParseFloat(zoom, 64); err == nil {
			if val > 0 {
				ctx.zoom = val
			}
		}
	}

	return ctx

}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			r := uint8(n * 1)
			g := uint8(n * 1)
			b := uint8(n * 11)
			return color.RGBA{R: r, G: g, B: b, A: 255}
		}
	}
	return color.Black
}
