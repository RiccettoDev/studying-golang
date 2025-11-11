package main

import (
	"fmt"
)

type client struct { // struct é uma estrutura com tipos diferentes(podem ser iguais) pré definidos
	name  string // digamos que se parece uma classe
	age   int
	smoke bool
}

func main() {
	client1 := client{
		name:  "José",
		age:   30,
		smoke: true,
	}

	client2 := client{
		name:  "Maria",
		age:   26,
		smoke: false,
	}

	fmt.Println(client1)
	fmt.Println(client2)
}
