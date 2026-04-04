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
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
			url = "http://" + url
		}
		go fetch(url, ch) // inicia uma gorrotina
	}
	for range os.Args[1:] {
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
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // evita vazamento de recursos
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
