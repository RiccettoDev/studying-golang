// Crie um loop utilizando a sintaxe: for condition {}
// Utilize-o para demonstrar os anos desde que você nasceu.

package main

import (
	"fmt"
	"strconv"
)

func main() {

	var x int

	for {
		fmt.Print("Digite o ano em que você nasceu: ")
		fmt.Scanln(&x)

		// Converte para string para contar dígitos
		s := strconv.Itoa(x)

		if len(s) == 4 && x < 2025 {
			break
		}

		fmt.Println("❌ Digite um ano válido de 4 dígitos e menor que 2025!")
	}

	for i := x; i <= 2025; i++ {
		fmt.Printf("Você viveu o ano %v\n", i)
	}

}
