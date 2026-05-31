/*Exercício 3.8: Renderizar fractais com níveis altos de zoom exige alta
precisão aritmética. Implemente o mesmo fractal usando quatro
representações numéricas diferentes: complex64, complex128, big.Float
e big.Rat. (Os dois últimos tipos encontram-se no pacote math/big.Float usa números de ponto flutuante quaisquer, porém com precisão
limitada; Rat usa números racionais com precisão ilimitada.) Como eles
se comparam quanto ao desempenho e ao uso de memória? Em que
níveis de zoom os artefatos de renderização tornam-se visíveis?*/

// Mandelbrot gera uma imagem PNG do fractal de Mandelbrot.
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/big"
	"math/cmplx"
	"os"
)

const (
	width, height = 1024, 1024
)

const (
	xcenter = -0.5
	ycenter = 0.0
	scale   = 4.0
)

var (
	xmin float64 = xcenter - scale/2
	xmax float64 = xcenter + scale/2
	ymin float64 = ycenter - scale/2
	ymax float64 = ycenter + scale/2
)

type BigComplex struct {
	Re *big.Float
	Im *big.Float
}

func main() {

	mode := mandelbrotModeFromArgs()

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		for px := 0; px < width; px++ {

			switch mode {
			case "complex128":
				z := pixelToComplex128(px, py)
				img.Set(px, py, mandelbrotComplex128(z))
			case "complex64":
				z := pixelToComplex64(px, py)
				img.Set(px, py, mandelbrotComplex64(z))

			}
			// Ponto (px, py) da imagem representa o valor complexo z

		}
	}
	png.Encode(os.Stdout, img) // NOTA: ignorando erros
}

func pixelToComplex128(px, py int) complex128 {
	x := float64(px)/width*(xmax-xmin) + xmin
	y := float64(py)/height*(ymax-ymin) + ymin
	return complex(x, y)
}

func mandelbrotComplex128(z complex128) color.RGBA {
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

// complex64

func pixelToComplex64(px, py int) complex64 {
	x := float32(px)/width*(float32(xmax)-float32(xmin)) + float32(xmin)
	y := float32(py)/height*(float32(ymax)-float32(ymin)) + float32(ymin)
	return complex(x, y)
}

func mandelbrotComplex64(z complex64) color.RGBA {
	const iterations = 200
	var v complex64

	for n := uint8(0); n < iterations; n++ {
		v = v*v + z

		if cmplx.Abs(complex128(v)) > 2 {
			r := uint8(n * 1)
			g := uint8(n * 1)
			b := uint8(n * 11)

			return color.RGBA{R: r, G: g, B: b, A: 255}
		}
	}

	return color.RGBA{R: 0, G: 0, B: 0, A: 255}
}

func mandelbrotModeFromArgs() string {
	//salva modo de execução a parti do arg
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Uso: go run main.go [complex128 | complex64]\n")
		os.Exit(1)
	}

	mode := os.Args[1]

	if mode != "complex128" && mode != "complex64" {
		fmt.Fprintf(os.Stderr, "Uso: go run main.go [complex128 | complex64]\n")
		os.Exit(2)
	}

	return mode
}
