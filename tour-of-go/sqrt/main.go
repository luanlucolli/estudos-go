package main

import (
	"fmt"
	"math"
)

func main() {
	x := float64(64)

	fmt.Printf("Função Sqrt personalizada: %f\n", Sqrt(x))
	fmt.Println("Função Sqrt da lib math:", math.Sqrt(x))
}

func Sqrt(x float64) float64 {
	fmt.Println("Calculando a raiz quadrada de:", x)

	if x < 0 || math.IsNaN(x) {
		return math.NaN()
	}

	if x == 0 || math.IsInf(x, 1) {
		return x
	}

	z := 1.0
	previousZ := z
	maxIterations := 20

	for range maxIterations {
		if math.IsInf(z, 0) {
			break
		}

		z -= (z*z - x) / (2 * z)
		if math.Abs(z-previousZ) < 1e-10 {
			break
		}
		previousZ = z
	}

	return z
}
