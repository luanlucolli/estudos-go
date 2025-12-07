package tour

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {

	fmt.Println("Calculando a raiz quadrada de:", x)

	if x < 0 {
		return math.NaN()
	}

	if x == 0 {
		return 0
	}

	z := 1.0
	previousZ := 0.0

	for {

		z -= (z*z - x) / (2 * z)
		if math.Abs(z-previousZ) < 1e-6 {
			break
		}
		previousZ = z
		fmt.Println(z)

	}

	return z
}
