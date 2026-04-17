/*
Exercício 2.1: Acrescente tipos, constantes e funções em tempconv para
processar temperaturas na escala Kelvin, em que zero Kelvin corresponde
a −273,15 °C e uma diferença de 1 K tem a mesma magnitude de 1 °C.
*/
package main

import (
	"fmt"

	"github.com/luanlucolli/estudos-go/a-linguagem-go/ch2/tempconv"
)

func main() {
	fmt.Println(tempconv.CToK(0))
	fmt.Println(tempconv.KToC(273.15))
}
