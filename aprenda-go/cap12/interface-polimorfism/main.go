package main

import (
	"fmt"
)

// -----------------------------------------------
// Definimos uma estrutura base chamada `pessoa`.
// Essa struct contém atributos comuns que podem
// ser compartilhados por diferentes tipos de pessoas.
// -----------------------------------------------
type pessoa struct {
	name    string
	surname string
	age     int
}

// -----------------------------------------------
// Agora criamos um novo tipo `dentista`, que "herda"
// (na prática, EMBUTE) a struct `pessoa`.
// Isso significa que `dentista` passa a ter acesso
// direto aos campos de `pessoa`, como `name`, `surname` e `age`.
//
// Além disso, `dentista` tem campos específicos,
// que só fazem sentido para um dentista.
// -----------------------------------------------
type dentista struct {
	pessoa
	dentesarrancados int
	salario          float64
}

// -----------------------------------------------
// Mesmo raciocínio: `arquiteto` também embute `pessoa`,
// mas adiciona seus próprios campos específicos.
// -----------------------------------------------
type arquiteto struct {
	pessoa
	tipodeconstrucao string
	tamanhodaloucura string
}

// -----------------------------------------------
// Aqui criamos um método com receptor do tipo `dentista`.
//
// Esse método "pertence" ao tipo `dentista`,
// e pode acessar todos os campos (inclusive os herdados de `pessoa`).
// -----------------------------------------------
func (x dentista) oibomdia() {
	fmt.Println("Meu nome é", x.name, "e eu já arranquei", x.dentesarrancados, "dentes... e ouve só: Bom dia!")
}

// -----------------------------------------------
// Outro método com o mesmo nome `oibomdia()`,
// mas agora pertencente ao tipo `arquiteto`.
//
// Isso é um exemplo de POLIMORFISMO EM AÇÃO:
// tipos diferentes (`arquiteto` e `dentista`) implementam
// o mesmo método com comportamentos diferentes.
// -----------------------------------------------
func (x arquiteto) oibomdia() {
	fmt.Println("Meu nome é", x.name, "e ouve só: Bom dia!")
}

// -----------------------------------------------
// Aqui definimos uma INTERFACE chamada `gente`.
// Interfaces em Go são coleções de comportamentos.
//
// A interface `gente` exige apenas um método: `oibomdia()`.
// Qualquer tipo que possua um método com essa assinatura
// automaticamente satisfaz a interface (não precisa declarar nada).
// -----------------------------------------------
type gente interface {
	oibomdia()
}

// -----------------------------------------------
// Função que recebe qualquer valor que satisfaça
// a interface `gente`.
//
// Isso é o ponto principal do POLIMORFISMO:
// a função `serhumano()` não precisa saber se o parâmetro
// é um `dentista`, `arquiteto`, ou outro tipo qualquer.
// Ela apenas exige que tenha o método `oibomdia()`.
// -----------------------------------------------
func serhumano(g gente) {
	g.oibomdia()
}

// -----------------------------------------------
// Função principal do programa.
// Aqui criamos instâncias concretas de `arquiteto` e `dentista`,
// inicializando seus campos (incluindo o campo embutido `pessoa`).
// -----------------------------------------------
func main() {
	// Cria um arquiteto, com uma struct `pessoa` dentro dele.
	teddyMosby := arquiteto{
		pessoa: pessoa{
			name:    "Teddy",
			surname: "Mosby",
			age:     28,
		},
		tipodeconstrucao: "prédios",
		tamanhodaloucura: "grande",
	}

	// Cria um dentista, também com uma `pessoa` embutida.
	barneyStinson := dentista{
		pessoa: pessoa{
			name:    "Barney",
			surname: "Stinson",
			age:     30,
		},
		dentesarrancados: 153,
		salario:          5000,
	}

	// Chamadas diretas aos métodos de cada tipo.
	// Cada um imprime seu próprio comportamento.
	teddyMosby.oibomdia()
	barneyStinson.oibomdia()

	// Chamadas polimórficas:
	// a função `serhumano()` recebe qualquer tipo
	// que satisfaça a interface `gente`.
	// Ou seja, `arquiteto` e `dentista` são "do tipo gente".
	serhumano(teddyMosby)
	serhumano(barneyStinson)
}
