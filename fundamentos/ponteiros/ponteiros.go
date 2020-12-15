package main

import "fmt"

func main() {
	x := 1

	var p *int = nil

	p = &x // pegando o endereço da variável
	*p++   // desreferenciando (pegando o valor)
	x++

	//Go não tem aritmética de ponteiros
	//p++

	fmt.Println(p, *p, x, &x)
}
