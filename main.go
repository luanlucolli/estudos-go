package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) float64 {

	fmt.Println("Calculando a raiz quadrada de:", x)

	if x < 0 {
		return math.NaN()
	}

	z := 1.0

	for math.Abs(z*z-x) > 1e-18 {

		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
	}

	return z
}

func main() {
	fmt.Printf("Função Sqrt personalizada: %f\n", sqrt(64))
	fmt.Println("Função Sqrt da lib math:", math.Sqrt(64))
}
