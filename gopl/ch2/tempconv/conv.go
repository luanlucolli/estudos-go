package tempconv

// CToF converte uma temperatura em Celsius para Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// CToK converte uma temperatura em Celsius para Kelvin.
func CToK(c Celsius) Kelvin { return Kelvin(c + 273.15) }

// FToC converte uma temperatura em Fahrenheit para Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// FToK converte uma temperatura em Fahrenheit para Kelvin.
func FToK(f Fahrenheit) Kelvin { return Kelvin((f-32)*5/9 + 273.15) }

// KToC converte uma temperatura em Kelvin para Celsius.
func KToC(k Kelvin) Celsius { return Celsius(k - 273.15) }

// KToF converte uma temperatura em Kelvin para Fahrenheit.
func KToF(k Kelvin) Fahrenheit { return Fahrenheit((k-273.15)*1.8 + 32) }
