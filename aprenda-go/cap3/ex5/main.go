// Utilizando a solução do exercício anterior:
//  1. Em package-level scope, utilizando a palavra-chave var, crie uma variável com o identificador "y". O tipo desta variável deve ser o tipo subjacente do tipo que você criou no exercício anterior.
//  2. Na função main:
//  1. Isto já deve estar feito:
//  1. Demonstre o valor da variável "x"
//  2. Demonstre o tipo da variável "x"
//  3. Atribua 42 à variável "x" utilizando o operador "="
//  4. Demonstre o valor da variável "x"
//  2. Agora faça tambem:
//  1. Utilize conversão para transformar o tipo do valor da variável "x" em seu tipo subjacente e, utilizando o operador "=", atribua o valor de "x" a "y"
//  2. Demonstre o valor de "y"
//  3. Demonstre o tipo de "y"
package main

import (
	"fmt"
)

type hotdog int

var x hotdog
var y int

func main() {

	fmt.Printf("Valor da variável x: %v\n", x)
	fmt.Printf("Tipo da variável x: %T\n", x)
	x = 42
	fmt.Printf("Valor da variável x: %v\n", x)
	fmt.Printf("Tipo da variável x: %T\n", x)

	y = int(x)
	fmt.Printf("Valor da variável y: %v\n", y)
	fmt.Printf("Tipo da variável y: %T\n", y)

}
