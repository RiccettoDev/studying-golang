package main // pacote principal

import (
	"fmt"
)

// Loops
// Laços de repetição
// Repetir tarefas

func main() {
	sum := 0
	// i++ soma 1
	// i-- subtrai 1
	for i := 0; i < 10; i++ {
		//fmt.Println("valor de i é: ", i)
		//fmt.Println("valor de sum é: ", sum)
		sum += i
		// é a mesma coisa que: sum = sum + 1
		// sum -= i -> é a mesma coisa que: sum = sum - i
		fmt.Println("valor de sum é: ", sum)
	}

	fmt.Println("=================")
	fmt.Println("Resultado: ", sum)

	// For Range
	frutas := []string{"laranja", "banana", "uva"}
	for index, fruta := range frutas {
		fmt.Println("Index: ", index)
		fmt.Println("Fruta: ", fruta)
	}

	fmt.Println("=================")

	for _, fruta := range frutas { // usamos o undescor caso não precise do index
		fmt.Println("Fruta: ", fruta)
	}

}
