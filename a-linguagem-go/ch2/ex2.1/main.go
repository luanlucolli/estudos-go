/*
Exercício 2.1: Acrescente tipos, constantes e funções em tempconv para
processar temperaturas na escala Kelvin, em que zero Kelvin corresponde
a −273,15 °C e uma diferença de 1 K tem a mesma magnitude de 1 °C.
*/

package main

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%g°K", k) }

// CToF converte uma temperatura em Celsius para Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// CToK converte uma temperatura em Celsius para Kelvin.
func CToK(c Celsius) Kelvin { return Kelvin(c + 273) }

// FToC converte uma temperatura em Fahrenheit para Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// FToK converte uma temperatura em Fahrenheit para Kelvin.
func FToK(f Fahrenheit) Kelvin { return Kelvin((f-32)*5/9 + 273) }

// KToC converte uma temperatura em Kelvin para Celsius.
func KToC(k Kelvin) Celsius { return Celsius(k - 273) }

// KToF converte uma temperatura em Kelvin para Fahrenheit.
func KToF(k Kelvin) Fahrenheit { return Fahrenheit((k-273)*1.8 + 32) }

func main() {
	fmt.Println(CToK(1))
	fmt.Println(KToC(1))
}
