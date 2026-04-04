/*Exercício 1.10: Encontre um site que gere uma grande quantidade de dados.
Investigue o caching executando fetchall duas vezes sucessivamente
para ver se o tempo informado sofre muita alteração. Você sempre obtém
o mesmo conteúdo? Modifique fetchall para exibir sua saída em um
arquivo para que ela possa ser examinada.*/

// Fetchall busca URLs em paralelo e informe os tempos gastos e os tamanhos.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
)

func main() {

	// valida numero minimo de args
	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "Uso: go run main.go exec[n] url1 url2...\n")
		os.Exit(1)
	}

	//valida arg
	if matched, _ := regexp.MatchString(`^exec\d+$`, os.Args[1]); !matched {
		fmt.Fprintf(os.Stderr, "Erro: use o padrão exec[n] (ex: exec1)\n")
		os.Exit(1)
	}

	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[2:] {
		if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
			url = "http://" + url
		}
		go fetch(url, ch) // inicia uma gorrotina
	}
	for range os.Args[2:] {
		fmt.Println(<-ch)
		// recebe do canal ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}
func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
		// envia para o canal ch
	}

	defer resp.Body.Close() // evita vazamento de recursos

	filename := strings.ReplaceAll(url, "://", "_")
	filename = strings.ReplaceAll(filename, "/", "_")
	filename = strings.ReplaceAll(filename, ":", "")

	file, err := os.Create(os.Args[1] + "_" + filename + ".txt")
	if err != nil {
		ch <- fmt.Sprintf("falha ao criar arquivo %s: %v", url, err)
		return
	}

	defer file.Close()

	nbytes, err := io.Copy(file, resp.Body)

	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
