package main

import (
	"fmt"
)

// structs embutidos são apenas structs dentro de structs

type pessoa struct {
	name string
	age  int
}

type profissional struct {
	pessoa
	titulo  string
	salario int
}

func main() {
	pessoa1 := pessoa{
		name: "José",
		age:  30,
	}

	profissional1 := profissional{
		pessoa:  pessoa1,
		titulo:  "padeiro",
		salario: 5000,
	}

	pessoa2 := pessoa{
		name: "Maria",
		age:  26,
	}

	profissional2 := profissional{
		pessoa:  pessoa2,
		titulo:  "manicure",
		salario: 7000,
	}

	fmt.Printf("%v de %v anos é %v ganhando %v \n", profissional1.name, profissional1.age, profissional1.titulo, profissional1.salario)
	fmt.Printf("%v de %v anos é %v ganhando %v \n", profissional2.name, profissional2.age, profissional2.titulo, profissional2.salario)
}

// estas são as saidas no terminal destes printf:

// José de 30 anos é padeiro ganhando 5000
// Maria de 26 anos é manicure ganhando 7000
