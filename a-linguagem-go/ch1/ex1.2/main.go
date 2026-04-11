/* Exercício 1.2: Modifique o programa echo para exibir o índice e o valor de
cada um de seus argumentos, um por linha. */

package main

import (
	"fmt"
	"os"
)

func main() {

	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("%d: %s\n", i, os.Args[i])
	}

}
