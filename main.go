package main

import (
	"fmt"

	"github.com/luanlucolli/hello-go/tour"
)

func main() {

	x := 1e50

	fmt.Printf("Função Sqrt personalizada: %f\n", tour.Sqrt(x))
	// fmt.Println("Função Sqrt da lib math:", math.Sqrt(64))

}
