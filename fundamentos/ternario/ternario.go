package main

import "fmt"

// GO NÂO TEM OPERADOR TERNARIO

func obterResultado(nota float64) string {
	if nota >= 6 {
		return "Aprovado"
	}
	return "Reprovado"
}

func main() {
	fmt.Println(obterResultado(6.3))
}
