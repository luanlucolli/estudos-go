/*Exercício 3.3: Pinte cada polígono de acordo com sua altura, de modo que os
picos tenham a cor vermelha ((#ff0000) e os vales sejam azuis (#0000ff).
*/

// Surface calcula uma renderização SVG de uma função de superfície 3D.
package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320
	// tamanho do canvas em pixels
	cells = 100
	// número de células da grade
	xyrange = 30.0
	// intervalos dos eixos (-xyrange..+xyrange)
	xyscale = width / 2 / xyrange // pixels por unidade x ou y
	zscale  = height * 0.4
	// pixels por unidade z
	angle = math.Pi / 6

// ângulo dos eixos x, y (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // seno(30°), cosseno(30°)

type surfaceFunc func(x, y float64) float64

func main() {

	funcaoEscolhida := sombrero

	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; strokewidth: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, zA, okA := corner(i+1, j, funcaoEscolhida)
			bx, by, zB, okB := corner(i, j, funcaoEscolhida)
			cx, cy, zC, okC := corner(i, j+1, funcaoEscolhida)
			dx, dy, zD, okD := corner(i+1, j+1, funcaoEscolhida)

			zAverage := (zA + zB + zC + zD) / 4.0

			if !okA || !okB || !okC || !okD {
				continue
			}

			color := polygonColor(zAverage)

			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='%s'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, color)
		}
	}

	fmt.Println("</svg>")
}
func corner(i, j int, f surfaceFunc) (float64, float64, float64, bool) {
	// Encontra o ponto (x,y) no canto da célula (i,j)
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	// Calcula a altura z da superfície
	z := f(x, y)
	if math.IsInf(z, 0) || math.IsNaN(z) {
		return 0, 0, 0, false
	}
	// Faz uma projeção isométrica de (x,y,z) sobre (sx,sy) do canvas SVG 2D
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z, true
}
func sombrero(x, y float64) float64 {
	r := math.Hypot(x, y) // distância de (0,0)
	return math.Sin(r) / r
}
func eggBox(x, y float64) float64 {
	return (math.Sin(x) + math.Sin(y)) * 0.1
}

func saddle(x, y float64) float64 {
	return (x*x - y*y) * 0.005
}

func moguls(x, y float64) float64 {
	return math.Cos(x) * math.Sin(y) * 0.1
}

func polygonColor(x float64) string {

	if x > 0 {
		return "#ff0000"
	}

	return "#0000ff"

}
