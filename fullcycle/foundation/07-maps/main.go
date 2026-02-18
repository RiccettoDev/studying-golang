package main

import (
	"fmt"
)

func main() {

	salarios := map[string]int{ // um map possui chave valor
		"padeiro":  5000,
		"vendedor": 4000,
		"Gerente":  7000,
	}

	for cargo, salario := range salarios { // está é uma forma diferente de fazer um for com range(Assim percorrendo todo o conteudo de "salarios")
		fmt.Printf("Cargo: %v, Salário: %v\n", cargo, salario)
	}

	delete(salarios, "padeiro") // desta maneira é possivel remover um elemento do maps atraves de sua chave

	fmt.Println("")
	fmt.Println("O padeiro foi demitido")
	fmt.Println("")

	for cargo, salario := range salarios {
		fmt.Printf("Cargo: %v, Salário: %v\n", cargo, salario)
	}

	salarios["novo padeiro"] = 5000 // desta maneira é possivel adicionar um novo elemento

	fmt.Println("")
	fmt.Println("Um novo padeiro foi contratado")
	fmt.Println("")

	for cargo, salario := range salarios {
		fmt.Printf("Cargo: %v, Salário: %v\n", cargo, salario)
	}

	for _, salario := range salarios { // usamos o _ para que não precisemos utilizar algum valor
		fmt.Printf("Salário: %v\n", salario)
	}

}
