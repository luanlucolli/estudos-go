/*
Exercício 2.5: A expressão x&(x-1) limpa o bit diferente de zero mais à
direita de x. Escreva uma versão de PopCount que conte bits usando esse
fato e avalie seu desempenho.
*/
package main

import "fmt"

func PopCount(x uint64) int {
	count := 0

	for x != 0 {
		x = x & (x - 1)
		count++
	}

	return count
}

func main() {
	var value uint64 = 22
	fmt.Println(PopCount(value))
}
