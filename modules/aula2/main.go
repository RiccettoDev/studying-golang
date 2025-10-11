package main // pacote principal

import ( // aqui a importação de outros pacotes
	"fmt"
)

func main() { // função principal

	fmt.Printf("Type: %T ---- Value: %v\n", true, true)
	fmt.Printf("Type: %T ---- Value: %v\n", "Eduardo", "Eduardo")
	fmt.Printf("Type: %T ---- Value: %v\n", 1, 1)
	fmt.Printf("Type: %T ---- Value: %v\n", "1", "1")
	fmt.Printf("Type: %T ---- Value: %v\n", 1.233, 1.233)

}

// Tipos
// bool (true/false)
// string
// int  int8  int16  int32  int64
// float (float64/float32) - decimal
