package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ { // essa é a estrutura básica do for em Golang
		fmt.Println(i) // em Go ; é um separador de instruções, onde aqui no for é necessário
	} // em outros lugares que comum mente utilizamos em outras linguagens não será necessário em Go

	fmt.Println("=========================================================")
	// OBS: Não existe WHILE em GOLANG
	// mas podemos fazer algo assim que funciona parecido com o WHILE:
	x := 0
	for x < 10 {
		fmt.Println(x)
		x++
	}

	fmt.Println("=========================================================")

	// assim pode ser declarado um loop infinito

	// for {
	// 	fmt.Println("infinit")
	// }

	cont := 0

	for {
		fmt.Printf("infinit %v\n", cont)
		cont++

		if cont == 10 {
			fmt.Println("Após exibir infinit 10 vezes a variável cont se tornou igual a 10")
			fmt.Println("Assim, entrando dentro do condicinal do nosso IF")
			fmt.Println("E ativando o 'break' que nos joga para fora deste loop infinito")
			break
		}
	}
}
