/* Exercício 1.4: Modifique dup2 para que exiba os nomes de todos os arquivos
em que cada linha duplicada ocorre. */

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	lineFiles := make(map[string]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, lineFiles, "stdin")
	} else {
		for _, arg := range files {
			fileName := arg
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, lineFiles, fileName)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
func countLines(f *os.File, counts map[string]int, lineFiles map[string]string, fileName string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	} // NOTA: ignorando erros em potencial de input.Err()
}
