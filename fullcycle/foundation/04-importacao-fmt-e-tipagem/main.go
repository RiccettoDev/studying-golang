package main

import (
	"fmt" // importe de um pacote de formatação
)

type ID int

func main() {

	var age ID = 10

	fmt.Printf("O tipo de age é %T e o valor é %v\n", age, age)
}
