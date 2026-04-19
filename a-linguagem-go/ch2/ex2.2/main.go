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
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/luanlucolli/estudos-go/a-linguagem-go/ch2/lenconv"
	"github.com/luanlucolli/estudos-go/a-linguagem-go/ch2/tempconv"
	"github.com/luanlucolli/estudos-go/a-linguagem-go/ch2/weightconv"
)

func main() {

	args := os.Args[1:]

	//checar se tem args
	if len(args) > 0 {
		for _, arg := range args {
			err := converteArg(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%v: erro na conversão: %v\n", arg, err)
			}
		}
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			err := converteArg(scanner.Text())
			if err != nil {
				fmt.Fprintf(os.Stderr, "%v: erro na conversão: %v\n", scanner.Text(), err)
			}
		}

		if err := scanner.Err(); err != nil {
			fmt.Fprintf(os.Stderr, "erro na leitura: %v\n", err)
			os.Exit(1)
		}
	}
}

func converteArg(s string) error {
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return fmt.Errorf("número inválido '%s': %v", s, err)
	}
	f := tempconv.Fahrenheit(v)
	c := tempconv.Celsius(v)
	ft := lenconv.Foot(v)
	mt := lenconv.Meter(v)
	p := weightconv.Pound(v)
	kg := weightconv.Kilogram(v)

	fmt.Printf("--- Conversões para %g ---\n", v)
	fmt.Printf("Temperatura: %s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
	fmt.Printf("Comprimento: %s = %s, %s = %s\n", ft, lenconv.FToM(ft), mt, lenconv.MToF(mt))
	fmt.Printf("Peso:        %s = %s, %s = %s\n\n", kg, weightconv.KToP(kg), p, weightconv.PToK(p))
	return nil
}
