package main

import (
	"fmt"
)

type Conta struct {
	saldo int
}

// Desta maneira apenas simula um saldo
func (c Conta) simular(valor int) int {
	c.saldo += valor
	fmt.Printf("Com a func simular a simulação do seu saldo é %v \n", c.saldo)
	return c.saldo
}

// Desta maneira altera seu saldo
func (c *Conta) simular2(valor int) int { // quando falamos com o * do ponteiro, não estamos mais tratando com uma cópia do valor e sim com o endereço de memória real
	c.saldo += valor
	fmt.Printf("Com a func simular2 a simulação do seu saldo é %v \n", c.saldo)
	return c.saldo
}

func main() {

	conta := Conta{saldo: 100}
	conta.simular(200)
	fmt.Printf("Mas seu saldo ainda é %v \n", conta.saldo)
	conta.simular2(400)
	fmt.Printf("Mas agora seu seu saldo é %v \n", conta.saldo)
}
