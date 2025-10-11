package main // pacote principal

import (
	"fmt"
)

// Maps Heterogêneos
// pode misturar tipos
// estrutura chave valor
// [key] = value
// chave tem um tipo e o valor pode ter outro
// map[k]v -> k = chave, v = valor

func main() {
	age := map[string]int{}
	age["Eduardo"] = 35
	age["Bento"] = 20
	fmt.Println(age)
	fmt.Println(age["Eduardo"])
	fmt.Println(age["Bento"])
}
