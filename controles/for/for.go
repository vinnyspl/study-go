package main

import (
	"fmt"
	"time"
)

func main() {

	i := 1

	for i <= 10 {
		fmt.Printf("%d", i)
		i++
	}

	for i := 0; i <= 20; i += 2 {
		fmt.Printf("%d", i)
		i++
	}

	fmt.Println("Misturando")

	for i := 0; i <= 15; i++ {
		if i%2 == 0 {
			fmt.Println("Par")
		} else {
			fmt.Println("Impar")
		}
	}

	for {
		//laço infinito
		fmt.Println("Ao inifinito e alem")
		time.Sleep(time.Second * 3)
	}

}
