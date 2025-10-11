package main // pacote principal

import (
	"fmt"
)

// Structs
// Forma de criar sua própria estrutura de dados
// Personalizar de acordo com a sua necessidade
// Podemos usar cários tipos diferentes

// type <nome da estrutura> struct { <campos> }
type Pessoa struct {
	Name    string
	Surname string
	Age     int
	Born    int
}

func main() {
	fmt.Println(Pessoa{"Eduardo", "Riccetto", 35, 1990})                           // retorno: {Eduardo Riccetto 35 1990}
	fmt.Println(Pessoa{Name: "Eduardo", Surname: "Riccetto", Age: 35, Born: 1990}) // retorno: {Eduardo Riccetto 35 1990}
	fmt.Println(Pessoa{Name: "Eduardo"})                                           // retorno: {Eduardo  0 0}

	p1 := Pessoa{Name: "Eduardo", Surname: "Riccetto", Age: 35, Born: 1990} // Atribuindo Pessoa a variável p1
	fmt.Println(p1.Surname)                                                 // retorno: Riccetto

	p2 := Pessoa{"Eduardo", "Riccetto", 35, 1990} // Atribuindo Pessoa a variável p2
	fmt.Println(p2.Born)                          //1990

	p2.Name = "José"
	fmt.Println(p2) // retorno: {José Riccetto 35 1990}
}
