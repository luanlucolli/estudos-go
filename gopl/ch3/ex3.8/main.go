/*
Exercício 3.8: Renderizar fractais com níveis altos de zoom exige alta
precisão aritmética. Implemente o mesmo fractal usando quatro
representações numéricas diferentes: complex64, complex128, big.Float
e big.Rat.

Uso:
	go run main.go complex64 > complex64.png
	go run main.go complex128 > complex128.png
	go run main.go bigfloat > bigfloat.png
	go run main.go bigrat > bigrat.png
*/

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
	width      = 1024
	height     = 1024
	iterations = 200
	prec       = 256
)

const (
	xcenter = -0.743643887037151
	ycenter = 0.131825904205330
	scale   = 1e-9
)

const (
	xmin = xcenter - scale/2
	xmax = xcenter + scale/2
	ymin = ycenter - scale/2
	ymax = ycenter + scale/2
)

// Valores usados por big.Float.
// Foram escritos como texto para não passar primeiro por float64
// e perder precisão antes do cálculo com big.Float.
var (
	bigFloatXCenter   = newBigFloat("-0.743643887037151")
	bigFloatYCenter   = newBigFloat("0.131825904205330")
	bigFloatHalfScale = newBigFloat("0.0000000005")
	bigFloatFour      = newBigFloat("4")

	bigFloatXMin = new(big.Float).SetPrec(prec).Sub(bigFloatXCenter, bigFloatHalfScale)
	bigFloatXMax = new(big.Float).SetPrec(prec).Add(bigFloatXCenter, bigFloatHalfScale)
	bigFloatYMin = new(big.Float).SetPrec(prec).Sub(bigFloatYCenter, bigFloatHalfScale)
	bigFloatYMax = new(big.Float).SetPrec(prec).Add(bigFloatYCenter, bigFloatHalfScale)

	bigFloatXSpan = new(big.Float).SetPrec(prec).Sub(bigFloatXMax, bigFloatXMin)
	bigFloatYSpan = new(big.Float).SetPrec(prec).Sub(bigFloatYMax, bigFloatYMin)
)

// Valores usados por big.Rat.
// Aqui o número é guardado como fração exata.
var (
	bigRatXCenter   = newBigRat("-0.743643887037151")
	bigRatYCenter   = newBigRat("0.131825904205330")
	bigRatHalfScale = newBigRat("1/2000000000")
	bigRatFour      = newBigRat("4")

	bigRatXMin = new(big.Rat).Sub(bigRatXCenter, bigRatHalfScale)
	bigRatXMax = new(big.Rat).Add(bigRatXCenter, bigRatHalfScale)
	bigRatYMin = new(big.Rat).Sub(bigRatYCenter, bigRatHalfScale)
	bigRatYMax = new(big.Rat).Add(bigRatYCenter, bigRatHalfScale)

	bigRatXSpan = new(big.Rat).Sub(bigRatXMax, bigRatXMin)
	bigRatYSpan = new(big.Rat).Sub(bigRatYMax, bigRatYMin)
)

func main() {
	mode := mandelbrotModeFromArgs()

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	// Variáveis reutilizadas na conversão dos pixels dos modos big.
	bigFloatRe := new(big.Float).SetPrec(prec)
	bigFloatIm := new(big.Float).SetPrec(prec)

	bigRatRe := new(big.Rat)
	bigRatIm := new(big.Rat)

	for py := 0; py < height; py++ {
		for px := 0; px < width; px++ {
			switch mode {
			case "complex64":
				z := pixelToComplex64(px, py)
				img.Set(px, py, mandelbrotComplex64(z))

			case "complex128":
				z := pixelToComplex128(px, py)
				img.Set(px, py, mandelbrotComplex128(z))

			case "bigfloat":
				pixelToBigFloat(px, py, bigFloatRe, bigFloatIm)
				img.Set(px, py, mandelbrotBigFloat(bigFloatRe, bigFloatIm))

			case "bigrat":
				pixelToBigRat(px, py, bigRatRe, bigRatIm)
				img.Set(px, py, mandelbrotBigRat(bigRatRe, bigRatIm))
			}
		}
	}

	png.Encode(os.Stdout, img) // NOTA: ignorando erros
}

func newBigFloat(text string) *big.Float {
	value, _ := new(big.Float).SetPrec(prec).SetString(text)
	return value
}

func newBigRat(text string) *big.Rat {
	value, _ := new(big.Rat).SetString(text)
	return value
}

func pixelToComplex128(px, py int) complex128 {
	x := float64(px)/width*(xmax-xmin) + xmin
	y := float64(py)/height*(ymax-ymin) + ymin

	return complex(x, y)
}

func pixelToComplex64(px, py int) complex64 {
	x := float32(px)/width*(float32(xmax)-float32(xmin)) + float32(xmin)
	y := float32(py)/height*(float32(ymax)-float32(ymin)) + float32(ymin)

	return complex(x, y)
}

func pixelToBigFloat(px, py int, x, y *big.Float) {
	pxFloat := new(big.Float).SetPrec(prec).SetInt64(int64(px))
	pyFloat := new(big.Float).SetPrec(prec).SetInt64(int64(py))

	widthFloat := new(big.Float).SetPrec(prec).SetInt64(int64(width))
	heightFloat := new(big.Float).SetPrec(prec).SetInt64(int64(height))

	x.Quo(pxFloat, widthFloat)
	x.Mul(x, bigFloatXSpan)
	x.Add(x, bigFloatXMin)

	y.Quo(pyFloat, heightFloat)
	y.Mul(y, bigFloatYSpan)
	y.Add(y, bigFloatYMin)
}

func pixelToBigRat(px, py int, x, y *big.Rat) {
	x.SetFrac64(int64(px), int64(width))
	x.Mul(x, bigRatXSpan)
	x.Add(x, bigRatXMin)

	y.SetFrac64(int64(py), int64(height))
	y.Mul(y, bigRatYSpan)
	y.Add(y, bigRatYMin)
}

func mandelbrotComplex128(z complex128) color.RGBA {
	var v complex128

	for n := 0; n < iterations; n++ {
		v = v*v + z

		if cmplx.Abs(v) > 2 {
			return colorForIteration(n)
		}
	}

	return color.RGBA{A: 255}
}

func mandelbrotComplex64(z complex64) color.RGBA {
	var v complex64

	for n := 0; n < iterations; n++ {
		v = v*v + z

		if cmplx.Abs(complex128(v)) > 2 {
			return colorForIteration(n)
		}
	}

	return color.RGBA{A: 255}
}

func mandelbrotBigFloat(zRe, zIm *big.Float) color.RGBA {
	vRe := new(big.Float).SetPrec(prec)
	vIm := new(big.Float).SetPrec(prec)

	nextRe := new(big.Float).SetPrec(prec)
	nextIm := new(big.Float).SetPrec(prec)

	temp1 := new(big.Float).SetPrec(prec)
	temp2 := new(big.Float).SetPrec(prec)
	sum := new(big.Float).SetPrec(prec)

	for n := 0; n < iterations; n++ {
		squareAddBigFloat(
			vRe, vIm,
			zRe, zIm,
			nextRe, nextIm,
			temp1, temp2,
		)

		vRe, nextRe = nextRe, vRe
		vIm, nextIm = nextIm, vIm

		if escapedBigFloat(vRe, vIm, temp1, temp2, sum) {
			return colorForIteration(n)
		}
	}

	return color.RGBA{A: 255}
}

// Calcula:
//
//	nextRe = vRe*vRe - vIm*vIm + zRe
//	nextIm = 2*vRe*vIm + zIm
func squareAddBigFloat(
	vRe, vIm *big.Float,
	zRe, zIm *big.Float,
	nextRe, nextIm *big.Float,
	temp1, temp2 *big.Float,
) {
	temp1.Mul(vRe, vRe)
	temp2.Mul(vIm, vIm)

	nextRe.Sub(temp1, temp2)
	nextRe.Add(nextRe, zRe)

	temp1.Mul(vRe, vIm)
	temp1.Add(temp1, temp1)

	nextIm.Add(temp1, zIm)
}

func escapedBigFloat(
	vRe, vIm *big.Float,
	temp1, temp2, sum *big.Float,
) bool {
	temp1.Mul(vRe, vRe)
	temp2.Mul(vIm, vIm)

	sum.Add(temp1, temp2)

	return sum.Cmp(bigFloatFour) > 0
}

func mandelbrotBigRat(zRe, zIm *big.Rat) color.RGBA {
	vRe := new(big.Rat)
	vIm := new(big.Rat)

	nextRe := new(big.Rat)
	nextIm := new(big.Rat)

	temp1 := new(big.Rat)
	temp2 := new(big.Rat)
	sum := new(big.Rat)

	for n := 0; n < iterations; n++ {
		squareAddBigRat(
			vRe, vIm,
			zRe, zIm,
			nextRe, nextIm,
			temp1, temp2,
		)

		vRe, nextRe = nextRe, vRe
		vIm, nextIm = nextIm, vIm

		if escapedBigRat(vRe, vIm, temp1, temp2, sum) {
			return colorForIteration(n)
		}
	}

	return color.RGBA{A: 255}
}

// Calcula:
//
//	nextRe = vRe*vRe - vIm*vIm + zRe
//	nextIm = 2*vRe*vIm + zIm
func squareAddBigRat(
	vRe, vIm *big.Rat,
	zRe, zIm *big.Rat,
	nextRe, nextIm *big.Rat,
	temp1, temp2 *big.Rat,
) {
	temp1.Mul(vRe, vRe)
	temp2.Mul(vIm, vIm)

	nextRe.Sub(temp1, temp2)
	nextRe.Add(nextRe, zRe)

	temp1.Mul(vRe, vIm)
	temp1.Add(temp1, temp1)

	nextIm.Add(temp1, zIm)
}

func escapedBigRat(
	vRe, vIm *big.Rat,
	temp1, temp2, sum *big.Rat,
) bool {
	temp1.Mul(vRe, vRe)
	temp2.Mul(vIm, vIm)

	sum.Add(temp1, temp2)

	return sum.Cmp(bigRatFour) > 0
}

func colorForIteration(n int) color.RGBA {
	return color.RGBA{
		R: uint8(n),
		G: uint8(n),
		B: uint8(n * 11),
		A: 255,
	}
}

func mandelbrotModeFromArgs() string {
	if len(os.Args) < 2 {
		fmt.Fprintf(
			os.Stderr,
			"Uso: go run main.go [complex64 | complex128 | bigfloat | bigrat]\n",
		)
		os.Exit(1)
	}

	mode := os.Args[1]

	if mode != "complex64" &&
		mode != "complex128" &&
		mode != "bigfloat" &&
		mode != "bigrat" {
		fmt.Fprintf(
			os.Stderr,
			"Uso: go run main.go [complex64 | complex128 | bigfloat | bigrat]\n",
		)
		os.Exit(2)
	}

	return mode
}
