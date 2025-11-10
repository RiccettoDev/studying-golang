package main

import (
	"fmt"
)

func main() {
	sabores := []string{"portuguesa", "marguerita", "mussarela", "milho", "peperoni"}

	fmt.Println("Nossos sabores de pizza:")
	fmt.Println("")
	for _, valores := range sabores { // usamos o _ quando não queros utilizar uma variável, para que o golang não reclame
		fmt.Printf("Sabor %v\n", valores)
	}

	fatia := sabores[0:2] // isso significa que vai incluir a posição 0 e 1 do slice

	fmt.Println("Fatia de sabores de pizza:")
	fmt.Println("")
	for _, valores := range fatia {
		fmt.Printf("Sabor %v\n", valores)
	}

	tamanhoSlice := len(sabores)

	fatia2 := sabores[2:tamanhoSlice] // isso significa que vai incluir a posição 2 até o ultimo elemento do slice

	// outros exemplos

	//fatia2 := sabores[2:] //isso aqui vai do 2 até o final
	//fatia2 := sabores[:2] //isso aqui vai do 0 até 2

	// sabores[:] assim são todos os elementos de sabores

	fmt.Println("Fatia de sabores de pizza:")
	fmt.Println("")
	for _, valores := range fatia2 {
		fmt.Printf("Sabor %v\n", valores)
	}

	// a remoção de um item de um slice é praticamente uma nova atribuição para o próprio slice

	fmt.Println("-------------")
	fmt.Println("abores de pizza")
	fmt.Println(sabores)

	sabores = append(sabores[:2], sabores[3:]...) // desta maneira deletamos o sabor mussarela do slice

	// os 3... pontinhos informam que é para desestruturar e pegar o conteudo indicado dentro do valor

	fmt.Println("novos abores de pizza")
	fmt.Println(sabores)
}
