package main

import (
	"fmt"
	"time"
)

func main() {

	t := time.Now()

	fmt.Println(t)

	switch {
	case t.Hour() < 12:
		fmt.Println("Bom dia!")
	case t.Hour() < 18:
		fmt.Println("Boa Tarde!")
	default:
		fmt.Println("Boa noite!")
	}

	fmt.Println(verificaNota(7.9))
}

func verificaNota(n float64) string {
	switch {
	case n >= 8 && n <= 10:
		return "SS"
	case n >= 6 && n < 8:
		return "A"
	case n >= 4 && n < 6:
		return "B"
	case n >= 2 && n < 4:
		return "C"
	default:
		return "D"
	}
}
