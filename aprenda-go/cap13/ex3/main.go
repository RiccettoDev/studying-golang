// Utilize a declaração defer de maneira que demonstre que sua execução só ocorre ao final do contexto ao qual ela pertence.

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Aqui sem defer")
	fmt.Println(1)
	fmt.Println(2)
	fmt.Println(3)
	fmt.Println(4)
	fmt.Println(5)

	fmt.Println("Aqui com defer no número 2")
	fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	fmt.Println(4)
	fmt.Println(5)

	fmt.Println("O que fez com que o número 2 fosse exibido como última instrução")

}
