package main

import "fmt"

func notaParaConceito(n float64) string {
	if n >= 9 && n <= 10 {
		return "A"
	} else if n >= 6 && n <= 8 {
		return "B"
	} else if n >= 3 && n <= 5 {
		return "C"
	} else {
		return "D"
	}
}

func main() {
	fmt.Println("Sua nota é", notaParaConceito(10))
}
