package main

import "fmt"

func main() {
	// homôgenea (mesmo tipo) e estatica (fixo)
	var notas [3]float64
	fmt.Println(notas)

	notas[0], notas[1], notas[2] = 10, 8.9, 111.5
	fmt.Println(notas)

	total := 0.0
	for i := 0; i < len(notas); i++ {
		total += notas[i]
	}
	fmt.Printf("Total de %.2f\n", total)

	media := total / float64(len(notas))
	fmt.Printf("Media de %.2f\n", media)
}
