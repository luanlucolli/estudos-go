/*Exercício 1.9: Modifique fetch para exibir também o código de status
HTTP encontrado em resp.Status.*/

// Fetch exibe o conteúdo encontrado em cada URL especificada.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] { // aproveitei a solução do exercício 1.8
		if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("%s: %s\n", url, resp.Status)

		_, err = io.Copy(os.Stdout, resp.Body) // aproveitei a solução do exercício 1.7
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}
