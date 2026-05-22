/*Exercício 3.4: Seguindo a abordagem do exemplo Lissajous na seção 1.7, crie
um servidor web que calcule superfícies e escreva dados SVG ao cliente.
O servidor deve definir o cabeçalho Content-Type assim:
w.Header().Set("Content-Type", "image/svg+xml")
(Esse passo não foi necessário no exemplo de Lissajous porque o servidor
usa métodos heurísticos padrão para reconhecer formatos comuns como
PNG a partir dos primeiros 512 bytes da resposta e gera o cabeçalho
apropriado.) Permita que o cliente especifique valores como altura,
largura e cor como parâmetros da requisição HTTP.*/

package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
)

// molde da função passada como argumento na função corner

func main() {

	// cada requisição chama surface3D
	http.HandleFunc("/surface", surface3D)

	// inicia o servidor
	log.Fatal(http.ListenAndServe("localhost:8000", nil))

}

func surface3D(w http.ResponseWriter, r *http.Request) {

	// verifica se o path é /favicon.ico e retorna caso seja
	if r.URL.Path == "/favicon.ico" {
		return
	}

	const (
		cells   = 100.0 // número de células da grade
		xyrange = 30.0  // intervalos dos eixos (-xyrange..+xyrange)
	)

	surfaceColors := map[string]string{
		"maxColor": "#ff0000",
		"minColor": "#0000ff",
	}

	surfaceValues := map[string]float64{
		"width":  600,
		"height": 320,
	}

	query := r.URL.Query()

	for key := range surfaceValues {

		valStr := query.Get(key)

		if valStr == "" {
			continue
		}

		valFloat, err := strconv.ParseFloat(valStr, 64)
		if err != nil {
			log.Printf("Erro ao converter query key '%s': %v", key, err)
			continue
		}

		surfaceValues[key] = valFloat

	}

	for key := range surfaceColors {

		valStr := query.Get(key)

		if valStr == "" {
			continue
		}

		surfaceColors[key] = valStr

	}

	width := surfaceValues["width"]
	height := surfaceValues["height"]
	maxColor := surfaceColors["maxColor"]
	minColor := surfaceColors["minColor"]
	xyscale := width / 2 / xyrange // pixels por unidade x ou y
	zscale := height * 0.4         // pixels por unidade z
	angle := math.Pi / 6
	sin30 := math.Sin(angle)
	cos30 := math.Cos(angle)

	type surfaceFunc func(x, y float64) float64

	corner := func(i, j int, f surfaceFunc) (float64, float64, float64, bool) {
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

	// query params deve aceitar altura, largura e cor

	w.Header().Set("Content-Type", "image/svg+xml")

	funcaoEscolhida := eggBox

	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; strokewidth: 0.7' "+
		"width='%g' height='%g'>", width, height)
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

			color := polygonColor(zAverage, maxColor, minColor)

			fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='%s'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, color)
		}
	}

	fmt.Fprintf(w, "</svg>")

}

/*
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
*/

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

func polygonColor(x float64, maxColor, minColor string) string {

	if x > 0 {
		return maxColor
	}

	return minColor

}
