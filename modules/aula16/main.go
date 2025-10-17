package main

import "net/http"

func main() {
	http.HandleFunc("/", HelloWord)   // rota principal chama a função HelloWord
	http.ListenAndServe(":8088", nil) // levanta um servidor na porta 8088, que pode ser acessada em http://localhost:8088/
}

func HelloWord(w http.ResponseWriter, r *http.Request) { // função HelloWord imprime na tela a mensagem: Hello Word Golang!
	w.Write([]byte("Hello Word Golang!"))
}
