/* Exercício 2.4: Escreva uma versão de PopCount que conte bits deslocando
seu argumento pelas 64 posições dos bits, testando o bit mais à direita a
cada vez. Compare seu desempenho com a versão que faz consultas na
tabela. */

package main

import "fmt"

// PopCount devolve a população (número de bits definidos) de x
func PopCount(x uint64) int {
	count := 0
	for i := 0; i < 64 && x != 0; i++ {
		count += int(x & 1)
		x >>= 1
	}
	return count
}

func main() {
	var value uint64 = 22

	fmt.Println(PopCount(value))
}
