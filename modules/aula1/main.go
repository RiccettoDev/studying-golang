package main // pacote principal

import ( // aqui a importação de outros pacotes
	print "fmt" // é possivel definir um apelido para os pacotes importados adicionando um nome antes do pacote
	"strings"
)

func main() { // função principal

	print.Println("Hello World") // aqui a função Println do pacote fmt --> printa no terminal um valor

	result := strings.Split("Eduardo", "") // salva o valor neste formato: [E d u a r d o]
	print.Println(result)

	result2 := strings.Split("Maça, Banana, Uva", ",") // salva o valor neste formato: [Maça Banana Uva]
	print.Println(result2)

}

// para compilar o código podemos usar:  "go build <nome e endereço do arquivo>"
