package main // pacote principal

import ( // aqui a importação de outros pacotes
	"fmt"
)

func main() { // função principal

	// variáveis

	// variável + nome da variável + tipo
	var nome string
	nome = "Eduardo"
	fmt.Println("Nome: ", nome)

	nome = "Eduardo Riccetto"
	fmt.Println("Nome alterado: ", nome)

	// não necessáriamente é preciso passar o tipo pois o gol é capaz de compreender o tipo
	var sobrenome = "Riccetto"
	fmt.Printf("Type: %T ---- Value: %v\n", sobrenome, sobrenome)

	//===============

	var x, y int = 1, 2
	fmt.Println(x)
	fmt.Println(y)

	//=============== outra forma de atribuir variável é com :=
	cachorro := "Pug"
	fmt.Println(cachorro)

	//===============

	var i int
	var f float64
	var b bool
	var s string

	fmt.Printf("Inteiro: %v\n", i)
	fmt.Printf("Float: %v\n", f)
	fmt.Printf("Bool: %v\n", b)
	fmt.Printf("String: %v\n", s)

	const idade = 18   // constante não pode ser alterada, por exemplo não podendo atribuir na sequencia //idade = 19
	fmt.Println(idade) // pq ela já foi atribuida como 18

}
