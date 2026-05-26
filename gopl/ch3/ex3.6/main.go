/*Exercício 3.6: Superamostragem (supersampling) é uma técnica para reduzir
o efeito de pixelation2, calculando o valor da cor em vários pontos em
cada pixel e tirando a média. O método mais simples é dividir cada pixel
em quatro “subpixels”. Implemente isso.*/

// Mandelbrot gera uma imagem PNG do fractal de Mandelbrot.
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

const (
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	width, height          = 1024, 1024
)

func main() {

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		for px := 0; px < width; px++ {

			// Ponto (px, py) da imagem representa o valor complexo z
			img.Set(px, py, supersample(px, py))

		}
	}
	png.Encode(os.Stdout, img) // NOTA: ignorando erros
}

func mandelbrot(z complex128) color.RGBA {
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

	return color.RGBA{R: 0, G: 0, B: 0, A: 255}
}

func pixelToComplex(px, py float64) complex128 {
	x := float64(px)/width*(xmax-xmin) + xmin
	y := float64(py)/height*(ymax-ymin) + ymin
	return complex(x, y)
}

func supersample(px, py int) color.Color {
	dx := []float64{0.25, 0.75}
	dy := []float64{0.25, 0.75}

	var somaR, somaG, somaB uint32

	for _, deslocx := range dx {
		for _, deslocy := range dy {
			z := pixelToComplex(float64(px)+deslocx, float64(py)+deslocy)
			cor := mandelbrot(z)

			somaR += uint32(cor.R)
			somaG += uint32(cor.G)
			somaB += uint32(cor.B)
		}
	}

	return color.RGBA{
		R: uint8(somaR / 4),
		G: uint8(somaG / 4),
		B: uint8(somaB / 4),
		A: 255,
	}
}
