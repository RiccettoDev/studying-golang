package main

import "fmt"

func main() {

	result, str := sum(1, 2)

	fmt.Println(result, str)

	resultSumAll := sumAll(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	fmt.Println(resultSumAll)

}

func sum(a, b int) (string, int) { // podemos ter mais de um tipo de retorno
	result := a + b

	return "O resultado da soma é: ", result
}

func sumAll(a ...int) int { // é possivel adicionar infititos parametros de valores
	result := 0

	for _, v := range a {
		result += v
	}

	return result
}
