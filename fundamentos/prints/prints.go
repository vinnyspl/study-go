package main

import "fmt"

func main() {
	fmt.Print("Nome")
	fmt.Print(" Sobre nome")

	fmt.Println("Nova")
	fmt.Println("linha.")

	x := 3.141516

	//fmt.Println("O valor de x é " + x)
	xs := fmt.Sprint(x)
	fmt.Println("O valor de x é" + xs)
	fmt.Println("O valor de x é", x)

	fmt.Println("O valor de x é %.2f.", x)

	a := 1
	b := 1.99999
	c := false
	d := "opa"
	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)
	fmt.Printf("\n%v %v %v %v", a, b, c, d)
}
