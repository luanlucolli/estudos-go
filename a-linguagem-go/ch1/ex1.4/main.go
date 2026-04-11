/* Exercício 1.4: Modifique dup2 para que exiba os nomes de todos os arquivos
em que cada linha duplicada ocorre. */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	lineFiles := make(map[string]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, lineFiles, "stdin")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, lineFiles, arg)
			f.Close()
		}
	}

	printDuplicates(counts, lineFiles)

}

func countLines(f *os.File, counts map[string]int, lineFiles map[string]string, fileName string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		counts[line]++
		if !strings.Contains(lineFiles[line], fileName) {
			lineFiles[line] += fileName + " "
		}
	} // NOTA: ignorando erros em potencial de input.Err()
}

func printDuplicates(counts map[string]int, lineFiles map[string]string) {
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%s\n", n, line, lineFiles[line])
		}
	}
}
