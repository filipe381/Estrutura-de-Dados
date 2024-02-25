package main

import (
	"fmt"
)

// func media aritmetica
func media(valores ...float64) float64 {
	if len(valores) == 0 {
		return 0
	}
	soma := 0.0
	for _, valor := range valores {
		soma += valor
	}

	media := soma / float64(len(valores))
	return media
}

// func paridade par ou impar
func paridade(numero int) string {
	if numero%2 == 0 {
		return "par"
	}
	return "impar"
}

// func potencia
func potencia(base, expoente int) int {
	resultado := 1
	for i := 0; i < expoente; i++ {
		resultado *= base
	}
	return resultado
}

// func celsius para fahrenheit
func fcelsius(celsius float64) float64 {
	return (celsius * 9 / 5) + 32
}

func main() {
	valores := []float64{2, 3, 4, 5, 6, 6, 7, 8}
	media := media(valores...)
	fmt.Printf("Média: %.2f\n", media)

	numero := 54998496
	fmt.Printf("%d é %s\n", numero, paridade(numero))

	base, expoente := 4, 9
	potencia := potencia(base, expoente)
	fmt.Printf("%d elevado a %d é %d\n", base, expoente, potencia)

	celsius := 25.0
	fahrenheit := fcelsius(celsius)
	fmt.Printf("%.1f celsius é igual a %.1f fahrenheit\n", celsius, fahrenheit)
}
