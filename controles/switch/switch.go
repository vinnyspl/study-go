package main

import "fmt"

func verificarNota(n float64) string {
	var nota = int(n)
	switch nota {
	case 10:
		fallthrough
	case 9:
		return "A"
	case 8, 7:
		return "B"
	case 6, 5:
		return "C"
	case 4, 3:
		return "D"
	case 2, 1, 0:
		return "E"
	default:
		return "Nota invalida!"
	}

}

func main() {
	fmt.Println(verificarNota(2.9))
	fmt.Println(verificarNota(10))
	fmt.Println(verificarNota(11))
}
