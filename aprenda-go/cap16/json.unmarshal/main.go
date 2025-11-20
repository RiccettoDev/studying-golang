// agora vamos transformar json em GO
// isto aqui é um json
// {"Name":"José","Age":30}

package main

import (
	"encoding/json"
	"fmt"
)

// aqui a informação que eu quero decodificar
type information struct {
	AquiONome string `json:"Name"` // isso aqui `json:"Name"` se chama encoder para caso o nome do seu atributo seja diferente do atributo em json
	Age       int    `json:"Age"`
}

func main() {

	resposta := []byte(`{"Name":"José","Age":30}`) // passamos um slice of bytes []byte(`{json}`)

	var dados information
	err := json.Unmarshal(resposta, &dados)
	// estamos passando a resposta, para o endereço de memória onde queremos salvar
	if err != nil {
		fmt.Println(err) // aqui apenas verificando se retorna erro
	}

	fmt.Println(dados)
	fmt.Println(dados.AquiONome)
	fmt.Println(dados.Age)
}

// respostas do terminal:
// {José 30}
// José
// 30
