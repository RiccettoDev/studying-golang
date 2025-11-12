package main

import (
	"fmt"
)

// defer significa adiar, postergar

// a ordem das instruções em uma função é importante e é executadade cima para baixo.

// caso alguma instrução possua defer, isso faz com que está instrução vá para o final da execução do programa.

func main() {
	fmt.Println("com defer, vai por ultimo")
	fmt.Println("sem defer, vem primeiro")
	fmt.Println("")

	defer fmt.Println("1")
	fmt.Println("2")
	fmt.Println("3")
	fmt.Println("4")

	// por exemplo aqui ao invez de exibir:
	// 1
	// 2
	// 3
	// 4

	// será exibido:
	// 2
	// 3
	// 4
	// 1

}
