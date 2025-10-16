package main

import "fmt"

func main() {

	a := 10

	fmt.Print(a)    // printou 10% -----> aqui apenas não pula linha
	fmt.Println("") // apenas para pular uma linha aqui
	fmt.Println(a)  // printou 10 -------> aqui mostrou o valor e pulou uma linha
	fmt.Println(&a) // printou 0x14000110020  --------> aqui mostrou o endereço de memória

	var ponteiro *int = &a // armazenou o endereço de memória

	fmt.Println(ponteiro)  // printou 0x14000110020 -----> aqui mostra o endereço de memória
	fmt.Println(*ponteiro) // printou 10 -----> o * indica para mostrar o valor no endereço de memória

	*ponteiro = 50

	fmt.Println(*ponteiro) // novo valor 50
	fmt.Println(ponteiro)  // no mesmo endereço de memória

	// e se vc parar para pensar a variavel "a" aponta para esse endereço de memória...
	// então se verificarmos o valor de "a" agora, ele deve ter sido alterado para 50

	fmt.Println("O valor de a agora é:", a) // printou 50
}
