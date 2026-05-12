package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Quando os nomes dos campos no JSON são diferentes dos nomes na struct, usamos tags.
// Tags são anotações entre crases após o tipo do campo. Exemplo:
//
//	Nome string `json:"name"`
//
// Isso faz com que o campo "Nome" da struct mapeie para "name" no JSON.
// Também podemos usar `json:"campo,omitempty"` para omitir o campo se estiver vazio,
// ou `json:"-"` para ignorar o campo completamente na serialização/deserialização.
type Conta struct {
	Numero int `json:"n"`
	Saldo  int `json:"s"`
}

func main() {
	conta := Conta{Numero: 1, Saldo: 100} // transformando uma struct em json
	res, err := json.Marshal(conta)
	if err != nil {
		panic(err)
	}
	println(string(res))

	conta2 := Conta{Numero: 2, Saldo: 200} // criamos o encoder, para quando não queremos salvar em uma variavel e apenas ler isso no terminal
	encoder := json.NewEncoder(os.Stdout)
	encoder.Encode(conta2)

	// agora vamos fazer ao contrario vamos reverter um json para struct
	jsonPuro := []byte(`{"n":3,"s":500}`) // mostrando que aqui poderia estar Numero ou n como esta. Falando sobre as tags no json
	var contaX Conta
	err = json.Unmarshal(jsonPuro, &contaX)
	if err != nil {
		println(err)
	}
	fmt.Printf("Conta numero %v com saldo de %v \n", contaX.Numero, contaX.Saldo)
}
