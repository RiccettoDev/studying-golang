package main

import (
	"fmt"
)

func main() {
	soma(3, 2)
	fmt.Println("==============")
	fmt.Println(somaInfinit(2, 3, 5))
}

func soma(x, y int) int {
	fmt.Println(x + y)
	return x + y
}

// é possivel inserir inumers parametros assim: (z ...int)
// é possivel definir inumeros retornos assim: (string, int, string, int)
// você pode ter quantos parametros quiser em uma função
// mas o elemento variádico deve ser o ultimo
// por exemplo vc não pode ter isso: (z ...int, s string)
// mas você pode ter isso: (s string, z ...int)
func somaInfinit(z ...int) (string, int, string, int) {
	soma := 0
	for _, v := range z {
		soma += v
	}

	return "Quantidade de numeros inseridos", len(z), "tem a soma de:", soma
}
