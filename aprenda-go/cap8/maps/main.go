package main

import (
	"fmt"
)

func main() {

	fmt.Println("==================")
	fmt.Println("Maps")
	fmt.Println("==================")

	agenda := map[string]int{ // um map é constituido por keys e values
		"Joao":   19999999999, // no caso aqui a keys é a string e o values é o int
		"Maria":  19999999998,
		"Carlos": 19999999997,
	}

	fmt.Println("")
	fmt.Println(agenda) // este é o retorno: map[Carlos:19999999997 Joao:19999999999 Maria:19999999998]

	// perceba que o retorno não ocorre de forma ordenada de acordo com a estrutura declarada
	// o retorno veio em ordem alfabética

	fmt.Println("")
	fmt.Println("Demonstrando somente um valor, por exemplo valor de Joao:")
	fmt.Println(agenda["Joao"])

	agenda["Eduardo"] = 19999999977

	fmt.Println("")
	fmt.Println("adicionando um novo Eduaro")
	fmt.Println(agenda["Eduardo"])

	exist, ok := agenda["fantasma"] // adicionamos o ok, fara verificar boolean
	// desta forma vemos que o valor 0 retornado não é um valor zero...
	// e sim um valor 0 representando false... o que significa que é inexistente, potanto este valor "fantasma" não existe

	fmt.Println("")
	fmt.Println("verificando um valor inexistente: 'fantasma'")
	fmt.Println(exist)
	fmt.Println(ok)
}
