// um pointeiro faz referência a uma localização na memória

package main

import (
	"fmt"
)

func main() {
	x := 10
	y := &x

	fmt.Printf("O valor de x é: %v\n", x)
	fmt.Printf("O valor de y é: %v\n", y)

	fmt.Println("")

	fmt.Printf("Com o * podemos ver o valor de uma memória\n")
	fmt.Printf("Onde *y agora é: %v\n", *y)
	fmt.Println("")

	fmt.Printf("O tipo de x é: %T \nO tipo de y é: %T\n", x, y)

}
