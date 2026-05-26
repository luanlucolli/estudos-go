/*Exercício 3.7: Outro fractal simples usa o método de Newton para encontrar
soluções complexas a uma função como z4−1 = 0. Sombreie cada ponto
de partida de acordo com o número de iterações necessárias para se
aproximar de uma das quatro raízes. Pinte cada ponto segundo a raiz da
qual ele se aproxima.*/

// gera uma imagem PNG do fractal de Newton.
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
			z := pixelToComplex(px, py)

			// Ponto (px, py) da imagem representa o valor complexo z
			img.Set(px, py, newton(z))

		}
	}
	png.Encode(os.Stdout, img) // NOTA: ignorando erros
}

func newton(z complex128) color.Color {
	const maxIter = 37
	const eps = 1e-6

	roots := []complex128{
		1 + 0i,
		-1 + 0i,
		0 + 1i,
		0 - 1i,
	}

	for n := 0; n < maxIter; n++ {
		den := 4 * z * z * z

		if cmplx.Abs(den) < 1e-12 {
			return color.Black
		}

		z = z - (z*z*z*z-1)/den

		var base color.RGBA
		var converged bool

		if cmplx.Abs(z-roots[0]) < eps {
			base = color.RGBA{R: 255, G: 0, B: 0, A: 255}
			converged = true
		} else if cmplx.Abs(z-roots[1]) < eps {
			base = color.RGBA{R: 0, G: 255, B: 0, A: 255}
			converged = true
		} else if cmplx.Abs(z-roots[2]) < eps {
			base = color.RGBA{R: 0, G: 0, B: 255, A: 255}
			converged = true
		} else if cmplx.Abs(z-roots[3]) < eps {
			base = color.RGBA{R: 255, G: 255, B: 0, A: 255}
			converged = true
		}

		if converged {
			shade := float64(maxIter-n) / float64(maxIter)

			return color.RGBA{
				R: uint8(float64(base.R) * shade),
				G: uint8(float64(base.G) * shade),
				B: uint8(float64(base.B) * shade),
				A: 255,
			}
		}
	}

	return color.Black
}

func pixelToComplex(px, py int) complex128 {
	x := float64(px)/width*(xmax-xmin) + xmin
	y := float64(py)/height*(ymax-ymin) + ymin
	return complex(x, y)
}

/*func supersample(px, py int) color.Color {
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
}*/
