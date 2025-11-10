package main

import (
	"fmt"
)

func main() {

	// tt := [][]int{
	// 	[]int{1, 2, 3},
	// 	[]int{4, 5, 6},
	// 	[]int{7, 8, 9},
	// }   // pode ser feito assim ou igual fizemos abaixo, da na mesma
	// mas assim o go vai reclamar que avisar o tipo está sendo redundante, pois o tipo já foi declarado no início

	ss := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println(ss[1][1]) // aqui vai exibir o 5 pois olhando em linha ele esta na posição 1 e em coluna na posição 1 também
	fmt.Println(ss[1][1]) // aqui vai exibir o 5 pois olhando em linha ele esta na posição 1 e em coluna na posição 1 também

	// funciona como se fosse o endereçamento e acesso de uma tabela em excel
}
