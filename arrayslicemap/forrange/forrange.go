package main

import "fmt"

func main() {
	numeros := [...]int{1, 2, 3, 4, 5, 6} // com os tres pontos o compilador conta o array

	for i, numero := range numeros {
		fmt.Printf("%d) %d\n", i, numero)
	}

	for _, num := range numeros {
		fmt.Printf("%d", num)
	}
}
