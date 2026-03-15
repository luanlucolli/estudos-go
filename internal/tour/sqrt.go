package tour

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {

	fmt.Println("Calculando a raiz quadrada de:", x)

	// tratamento de casos de borda
	if x < 0 || math.IsNaN(x) {
		return math.NaN()
	}

	if x == 0 || math.IsInf(x, 1) {
		return x
	}

	// variaveis iniciais
	z := 1.0
	previousZ := z
	maxIterations := 20
	// loop que itera até z * z estiver próximo de x
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

	// retorna o valor aproximado da raiz quadrada
	return z
}
