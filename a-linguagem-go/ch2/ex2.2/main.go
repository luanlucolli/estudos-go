/*
Exercício 2.2: Escreva um programa de conversão de unidades de propósito
geral, análogo ao cf, que leia números de seus argumentos de linha de
comando ou da entrada-padrão se não houver argumentos, e converta
cada número em unidades como temperatura em Celsius e em
Fahrenheit, comprimento em pés e metros, peso em libras e quilogramas
e operações semelhantes.
*/
package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/luanlucolli/estudos-go/a-linguagem-go/ch2/lenconv"
	"github.com/luanlucolli/estudos-go/a-linguagem-go/ch2/tempconv"
)

func main() {
	//fmt.Println(lenconv.TempFoot.String())
	//fmt.Println(lenconv.TempMeter.String())

	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}

		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		ft := lenconv.Foot(t)
		mt := lenconv.Meter(t)
		//converter pra pés 1 pé = 0,3048 metros
		//converter pra metros 1 metro = 3,28084 pés
		//peso em libras
		//peso em quilogramas
		fmt.Printf("%s = %s, %s = %s\n%s = %s, %s = %s\n",
			f, tempconv.FToC(f), c, tempconv.CToF(c), ft, lenconv.FToM(ft), mt, lenconv.MToF(mt))
	}
}
