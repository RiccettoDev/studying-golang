package main // pacote principal

import ( // aqui a importação de outros pacotes
	"fmt"
)

// funções

func main() { // função principal
	x, y := 1, 2
	result := sum(x, y)
	fmt.Printf("A soma de %d + %d é igual a: %d\n", x, y, result)

	fmt.Println(printaNome("Eduardo")) // será printado o nome Eduardo duas vezes
}

func sum(x int, y int) int { // em uma função é preciso declarar os tipos dos parametros e tambem da saida que no caso é o int de fora dos parenteses
	sumXY := x + y
	return sumXY
}

func printaNome(nome string) (string, string) { // ao adicionar dois parametros de retorno, mais de um tipo de tipagem devem ficar salvo dentro deste segundo parenteses
	return nome, nome
}

// se uma função é criada com nome minusculo ela é privada deste pacote
// se uma função é criada com nome maiusculo ela é publica para ser usada em outros pacotes
