package main

import (
	"fmt"
	"go-api/modules/aula12/calc"
)

func main() {
	fmt.Println("Olá Mundo!")
	fmt.Println("Vamos somar")
	fmt.Println("-------------")

	result := calc.Sum(1, 2)
	fmt.Println("O resultado da soma é: ", result)

	fmt.Println("É possivel chamar uma variável de dentro de outro pacote: ", calc.DentroDoPacoteCalc)
}
