package main

import (
	"fmt"
	"sort"
)

func main() {
	x := []string{"d", "f", "b", "l", "a", "r"}
	y := []int{22, 10, 35, 102, 3}

	fmt.Println(x)
	sort.Strings(x) // ordenou o slice em ordem alfabética
	fmt.Println(x)

	fmt.Println(y)
	sort.Ints(y) // ordenou o slice em ordem crescente
	fmt.Println(y)
}
