package main

func main() {
	println(sum(1, 2, 3, 4, 5, 6, 7, 8, 9))
}

func sum(numeros ...int) int { // fazemos assim quando não sabemos a quantidade exata de parametros e assim podemos inserir diversos
	total := 0

	for _, valor := range numeros {
		total += valor
	}

	return total
}
