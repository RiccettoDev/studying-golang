package main

func somaInt(m map[string]int) int { // para somar inteiros
	var somaInt int
	for _, v := range m {
		somaInt += v
	}
	return somaInt
}

func somaFloat(m map[string]float64) float64 { // para somar flutuantes
	var somaFloat float64
	for _, v := range m {
		somaFloat += v
	}
	return somaFloat
}

// para que não seja necessário fazer duas funções iguais apenas por questões de tipos, usamos o generics
func soma[T int | float64](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func main() {
	m := map[string]int{"Edu": 1000, "João": 2000, "Maria": 3000}
	println(somaInt(m))
	m2 := map[string]float64{"Edu": 1000.2, "João": 2000.3, "Maria": 3000.0}
	println(somaFloat(m2))

	// Desta forma funcionamos apenas com uma unica função generica
	mt1 := map[string]int{"Edu": 1000, "João": 2000, "Maria": 3000}
	println(soma(mt1))
	mt2 := map[string]float64{"Edu": 1000.2, "João": 2000.3, "Maria": 3000.0}
	println(soma(mt2))
}
