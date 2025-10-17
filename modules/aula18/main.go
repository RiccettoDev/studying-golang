package main

import "fmt"

func SumGeneric[T int64 | float64](m map[string]T) T { // assim trabalhando com dois tipos diferentes na função
	var s T
	for _, v := range m {
		s += v
	}
	return s
}

// podendo tambem ser da maneira abaixo

type Number interface {
	int | int64 | float64
}

func SumGeneric2[T Number](m map[string]T) T { // assim trabalhando com dois tipos diferentes na função
	var s T
	for _, v := range m {
		s += v
	}
	return s
}

func main() {
	fmt.Println(SumGeneric(map[string]int64{"a": 1, "b": 2, "c": 3}))
	fmt.Println(SumGeneric(map[string]float64{"a": 1.1, "b": 2.22, "c": 3.333}))

	fmt.Println(SumGeneric2(map[string]int{"a": 1, "b": 2, "c": 3}))
	fmt.Println(SumGeneric2(map[string]int64{"a": 1, "b": 2, "c": 3}))
	fmt.Println(SumGeneric2(map[string]float64{"a": 1.1, "b": 2.22, "c": 3.333}))
}
