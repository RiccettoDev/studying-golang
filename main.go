package main // pacote principal

import ( // aqui a importação de outros pacotes
	"fmt"
	"strings"
)

func main() { // função principal
	fmt.Println("Hello World") // aqui a função Println do pacote fmt --> printa no terminal um valor

	result := strings.Split("Eduardo", "") // salva o valor neste formato: [E d u a r d o]
	fmt.Println(result)

	result2 := strings.Split("Maça, Banana, Uva", ",") // salva o valor neste formato: [Maça Banana Uva]
	fmt.Println(result2)

}
