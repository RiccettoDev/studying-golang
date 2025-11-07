package main

import (
	"fmt"
)

func main() {
	for i := 0; i <= 10; i++ { // aqui um loopç padrão exibindo de números de 0 a 10
		fmt.Println(i)
	}

	fmt.Println("==========================")

	for i := 0; i <= 10; i++ { // aqui um loopç padrão exibindo de números de 0 a 10
		if i%2 == 0 { // aqui um IF condicional verificando quando I divisivel por 2 for igaul a zero
			fmt.Println(i) // portanto exbindo somente numeros pares
		}
	}

	fmt.Println("==========================")

	for i := 0; i <= 10; i++ {
		if i%2 != 0 {
			fmt.Printf("%v é um número impar então vamos finalizar\n", i)
			break // aqui estamos verificando que um número é impar e estamos utilizando o break para finalizar a interação FOR
		}
		fmt.Println(i)
	}

	fmt.Println("==========================")

	for i := 0; i <= 10; i++ {
		if i%2 == 0 { // agora aqui estamos verificando quais os números que pares
			continue // e então usando o continue para pular esta interação caso a condição do numero par for verdadeira
		}
		fmt.Println(i) // e aqui exibindo os números impares
	}

	// então usamos o break para acabar com a interação FOR
	// e usamos o continue para acaber apenas um uma interação especifica...
	// digamos então acabar com um ciclo de loop especifico do FOR
}
