// conversão binário e hexadecimal

package main

import (
	"fmt"
)

func main() {
	a := 100

	fmt.Printf("Valores em diferentes formatos:\n")
	fmt.Printf("\n")
	fmt.Printf("Valor de 100 em decimal: %d\n", a)
	fmt.Printf("Valor de 100 em binário: t%b\n", a)
	fmt.Printf("Valor de 100 em hexadecimal: t%#x\n", a) // pode ser utilizado sem o #
	// o # serve apenas para que o valor seja exido com um prefixo...
	// por exemplo o numero 10 em hexadecimal equivale a letra "a"
	// estes são os casoa de exibição para o valor de 10
	// %v ---> decimal = 10
	// %b ---> binário = 1010
	// %x ---> valor em hexadecimal = a
	// %X ---> valor em hexadecimal = A  // aqui mostrando a direfença entre usar o x minusculo e o X maiusculo
	// %#X ---> valor em hexadecimal = 0xa   // usamos o # para adicionar o prefixo
	// %#X ---> valor em hexadecimal = 0xA  // aqui mostrando a direfença entre usar o x minusculo e o X maiusculo

	// e para que serve a utilização do prefixo?
	// utilizamos para como no exemplo, não confundirmos o valor "A" que é o valor 10 em hexadecimal
	// com o valor de "A" string, ou com algum "A" de variável.
}
