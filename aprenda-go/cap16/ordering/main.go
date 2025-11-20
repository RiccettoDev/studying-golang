package main

import (
	"fmt"
	"sort"
)

type pessoa struct {
	name        string
	age         int
	zodiacSigns string
}

type ordenarPorMenorAge []pessoa // criamos um tipo

// criamos métodos para este tipo
func (x ordenarPorMenorAge) Len() int {
	return len(x)
}

// criamos métodos para este tipo
func (x ordenarPorMenorAge) Less(i, j int) bool { // i e j seignificam indice e indice seguinte
	return x[i].age < x[j].age
}

// criamos métodos para este tipo
func (x ordenarPorMenorAge) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

// e com estes 3 métodos em nosso tipo, podemos utilizar a função sort.Sort

type ordenarPorMaiorAge []pessoa // criamos um novo tipo
func (x ordenarPorMaiorAge) Len() int {
	return len(x)
}
func (x ordenarPorMaiorAge) Less(i, j int) bool {
	return x[i].age > x[j].age // mudamos para uma ordenação decescente
}
func (x ordenarPorMaiorAge) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

func main() {
	listaDePessoas := []pessoa{
		{"João", 20, "Peixes"},
		{"Maria", 25, "Leão"},
		{"Claudia", 30, "Libra"},
		{"Felipe", 22, "Capricórnio"},
	}

	fmt.Println(listaDePessoas)                   // aqui mostramos na ordem em que fois cadastrado
	sort.Sort(ordenarPorMenorAge(listaDePessoas)) // aqui fazemos o sort para ordenar por menor idade
	fmt.Println(listaDePessoas)                   // e aqui mostramos a lista já ordenada por menor idade

	sort.Sort(ordenarPorMaiorAge(listaDePessoas)) // aqui fazemos o sort para ordenar por maior idade
	fmt.Println(listaDePessoas)                   // e aqui mostramos a lista já ordenada por maior idade
}

// saídas do terminal:
// [{João 20 Peixes} {Maria 25 Leão} {Claudia 30 Libra} {Felipe 22 Capricórnio}]
// [{João 20 Peixes} {Felipe 22 Capricórnio} {Maria 25 Leão} {Claudia 30 Libra}]
