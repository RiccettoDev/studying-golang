package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// manipulação de arquivos

	// criação e escrita
	f, err := os.Create("arquivo.txt") // cria um arquivo
	if err != nil {
		panic(err)
	}

	tamanho, err := f.WriteString("Hello World!\n") // escreve em um arquivo
	if err != nil {
		panic(err)
	}
	fmt.Printf("Arquivo criado com sucesso! Tamanho: %d bytes \n", tamanho)

	tamanho2, err := f.Write([]byte("Escrevendo dados no arquivo")) // escreve em um arquivo independente se é string ou não
	if err != nil {
		panic(err)
	}
	fmt.Printf("Escrevendo dados com sucesso! Tamanho: %d bytes\n", tamanho2)

	defer f.Close()

	// leitura de um arquivo

	arquivo, err := os.ReadFile("arquivo.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(arquivo))

	// leitura de pouco em pouco abrindo arquivo
	arquivo2, err := os.Open("arquivo.txt")
	if err != nil {
		panic(err)
	}
	println("Arquivo está aberto")
	reader := bufio.NewReader(arquivo2)
	buffer := make([]byte, 5)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}

}
