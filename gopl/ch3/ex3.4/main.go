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

type RenderContext struct {
	width, height      float64
	xyscale, zscale    float64
	sin30, cos30       float64
	cells              float64
	xyrange            float64
	angle              float64
	maxColor, minColor string
	surfacaceFunction  surfaceFunc
}

// molde da função passada como argumento na função corner
type surfaceFunc func(x, y float64) float64

func main() {

	// cada requisição chama surface3D
	http.HandleFunc("/surface", surface3D)

	// inicia o servidor
	log.Fatal(http.ListenAndServe("localhost:8000", nil))

}

func surface3D(w http.ResponseWriter, r *http.Request) {

	// define a forma da superficie 3d
	var surfaceFunction surfaceFunc = sombrero

	// cria maps com os valores aceitos no query params da url
	surfaceColors := map[string]string{
		"maxColor": "#ff0000",
		"minColor": "#0000ff",
	}
	surfaceValues := map[string]float64{
		"width":  600,
		"height": 320,
	}

	// chama a função que extrai os valores dos query params e atualiza nos maps
	processQueryParams(surfaceValues, surfaceColors, &surfaceFunction, r)

	// monta a struct e define valores que não dependem de operações com outros valores da mesma struct
	ctx := &RenderContext{
		width:    surfaceValues["width"],
		height:   surfaceValues["height"],
		maxColor: surfaceColors["maxColor"],
		minColor: surfaceColors["minColor"],
		cells:    100.0,
		xyrange:  30.0,
		angle:    math.Pi / 6,
	}

	// termina de definir valores da sruct
	ctx.xyscale = ctx.width / 2 / ctx.xyrange
	ctx.zscale = ctx.height * 0.4
	ctx.sin30 = math.Sin(ctx.angle)
	ctx.cos30 = math.Cos(ctx.angle)

	// seta o header para image/svg+xml
	w.Header().Set("Content-Type", "image/svg+xml")

	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; strokewidth: 0.7' "+
		"width='%g' height='%g'>", ctx.width, ctx.height)
	for i := 0; i < int(ctx.cells); i++ {
		for j := 0; j < int(ctx.cells); j++ {
			ax, ay, zA, okA := ctx.corner(i+1, j, surfaceFunction)
			bx, by, zB, okB := ctx.corner(i, j, surfaceFunction)
			cx, cy, zC, okC := ctx.corner(i, j+1, surfaceFunction)
			dx, dy, zD, okD := ctx.corner(i+1, j+1, surfaceFunction)

			zAverage := (zA + zB + zC + zD) / 4.0

			if !okA || !okB || !okC || !okD {
				continue
			}

			color := polygonColor(zAverage, ctx.maxColor, ctx.minColor)

			fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='%s'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, color)
		}
	}

	fmt.Fprintf(w, "</svg>")

}

// extrai os query params e atualiza os maps se houver valor nas keys
// não retorna nada, atualiza o map diretamente
func processQueryParams(surfaceValues map[string]float64, surfaceColors map[string]string, surfaceFunction *surfaceFunc, r *http.Request) {

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

	tipo := query.Get("type")

	switch tipo {
	case "sombrero":
		*surfaceFunction = sombrero
	case "eggbox":
		*surfaceFunction = eggBox
	case "saddle":
		*surfaceFunction = saddle
	case "moguls":
		*surfaceFunction = moguls
	}

}

func (ctx *RenderContext) corner(i, j int, f surfaceFunc) (float64, float64, float64, bool) { // Encontra o ponto (x,y) no canto da célula (i,j)
	x := ctx.xyrange * (float64(i)/ctx.cells - 0.5)
	y := ctx.xyrange * (float64(j)/ctx.cells - 0.5)
	// Calcula a altura z da superfície
	z := f(x, y)
	if math.IsInf(z, 0) || math.IsNaN(z) {
		return 0, 0, 0, false
	}
	// Faz uma projeção isométrica de (x,y,z) sobre (sx,sy) do canvas SVG 2D
	sx := ctx.width/2 + (x-y)*ctx.cos30*ctx.xyscale
	sy := ctx.height/2 + (x+y)*ctx.sin30*ctx.xyscale - z*ctx.zscale
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

func polygonColor(x float64, maxColor, minColor string) string {

	if x > 0 {
		return maxColor
	}

	return minColor

}
