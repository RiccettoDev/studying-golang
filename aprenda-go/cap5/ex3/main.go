// Crie constantes tipadas e não-tipadas.
// Demonstre seus valores.

package main

import (
	"fmt"
)

// Constante TIPADA - tipo definido na declaração
const a int = 10

// Constante NÃO-TIPADA - sem tipo na declaração
const b = 10

func main() {
	// Demonstração do conceito fundamental
	fmt.Println("=== CONCEITO FUNDAMENTAL ===")
	fmt.Printf("Constante TIPADA 'a':\n")
	fmt.Printf("  Tem tipo DEFINIDO na declaração: const a int = 10\n")
	fmt.Printf("  Tipo: %T\n", a)
	fmt.Printf("  Valor: %d\n", a)

	fmt.Printf("\nConstante NÃO-TIPADA 'b':\n")
	fmt.Printf("  NÃO tem tipo na declaração: const b = 10\n")
	fmt.Printf("  Tipo ATUAL: %T\n", b)
	fmt.Printf("  Valor: %d\n", b)
}
