/* Exercício 2.3: Reescreva PopCount para que use um loop no lugar de uma
expressão única. Compare o desempenho das duas versões. (A seção 11.4
mostra como comparar o desempenho de diferentes implementações de
forma sistemática.) */

package main

import "fmt"

// pc[i] é a população de i
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount devolve a população (número de bits definidos) de x
func PopCount(x uint64) int {
	result := 0
	for i := 0; i < 8; i++ {
		result += int(pc[byte(x>>(i*8))])
	}
	return result
}

func main() {
	var value uint64 = 22

	fmt.Println(PopCount(value))
}
