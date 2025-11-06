// Crie um programa que:
// Atribua um valor int a uma variável
// Demonstre este valor em decimal, binário e hexadecimal
// Desloque os bits dessa variável 1 para a esquerda, e atribua o resultado a outra variável
// Demonstre esta outra variável em decimal, binário e hexadecimal

package main

import (
	"fmt"
)

func main() {

	var number int

	fmt.Println("Digite um número inteiro")
	fmt.Scanln(&number)

	fmt.Printf("O valor digitado em Decimal é: %v\n", number)
	fmt.Printf("O valor digitado em Binário é: %b\n", number)
	fmt.Printf("O valor digitado em Hexadecimal é: %#X\n", number)

	var binario = fmt.Sprintf("%b", number)

	var deslocada = number << 1

	fmt.Printf("O valor de %v deslocado << 1 é igual a %b\n", binario, deslocada)

	fmt.Printf("O valor deslocado em Decimal é: %v\n", deslocada)
	fmt.Printf("O valor deslocado em Binário é: %b\n", deslocada)
	fmt.Printf("O valor deslocado em Hexadecimal é: %#X\n", deslocada)

}
