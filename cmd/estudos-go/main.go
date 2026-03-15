package main

import (
	"fmt"
	"math"

	"github.com/luanlucolli/estudos-go/internal/tour"
)

func main() {

	x := float64(1e307)

	fmt.Printf("Função Sqrt personalizada: %f\n", tour.Sqrt(x))
	fmt.Println("Função Sqrt da lib math:", math.Sqrt(64))

}
