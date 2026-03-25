/* Exercício 1.1: Modifique o programa echo para exibir também os.Args[0],
que é o nome do comando que o chamou. */

package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
