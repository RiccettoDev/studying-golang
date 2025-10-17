package main

import (
	"fmt"
	"sync"
	"time"
)

func cont(qtd int) { // esta função faz uma contagem exibindo os números como tipo um timer.
	for i := range qtd {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
} // então por exemplo ao chamar a função cont(10), ela vai durar 10 seguntos, exibindo: 0, 1, 2, 3, 4, 5, 6, 7, 8 e 9

func main() {
	var wg sync.WaitGroup

	wg.Add(2) // Diz que vamos esperar por 2 goroutines
	cont(10)
	cont(10)

	go cont(10) // estas duas chamadas acima vão durar 20 segundos, pois ocorre uma após a outra
	go cont(10) // ao adicionarmo a palavra go antes da função aqui, trabalhamos com mult threads
	// portanto as 3 funções rodando não vão durar agora 30 segundos para terminar

	wg.Wait() // Espera até que todas as goroutines terminem
	fmt.Println("Todas as goroutines terminaram!")
}
