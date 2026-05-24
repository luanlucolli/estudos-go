/*Exercício 3.5: Implemente o conjunto de Mandelbrot todo colorido usando a
função image.NewRGBA e o tipo color.RGBA ou color.YCbCr.*/

// Mandelbrot gera uma imagem PNG do fractal de Mandelbrot.
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)
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
	png.Encode(os.Stdout, img) // NOTA: ignorando erros
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
