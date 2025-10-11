package main // pacote principal

import (
	"fmt"
	"sort"
	"strings"
)

/*
=== ARRAYS E SLICES EM GO ===

ARRAYS:
- Estruturas de dados homogêneas (todos elementos do mesmo tipo)
- Tamanho fixo definido na declaração
- Sintaxe: var nome [tamanho]tipo
- Exemplo: [5]int{1, 2, 3, 4, 5}
- Arrays são valores, não referências
- Quando passados para funções, são copiados

SLICES:
- Arrays dinâmicos, mais flexíveis que arrays
- Tamanho pode variar durante execução
- Sintaxe: var nome []tipo ou make([]tipo, tamanho)
- Exemplo: []int{1, 2, 3, 4, 5}
- Slices são referências para arrays subjacentes
- Mais comumente usados que arrays em Go

OPERAÇÕES COMUNS:
- len(): retorna o comprimento
- cap(): retorna a capacidade (apenas slices)
- append(): adiciona elementos (apenas slices)
- copy(): copia elementos entre slices
*/

/*
=== MAPS EM GO ===

DEFINIÇÃO:
- Estrutura de dados chave-valor (key-value)
- Chaves devem ser de tipos comparáveis (string, int, bool, arrays, structs sem slices/maps/funções)
- Valores podem ser de qualquer tipo
- Sintaxe: map[TipoChave]TipoValor
- Exemplo: map[string]int{"Eduardo": 34, "Bento": 20}

CARACTERÍSTICAS:
- Maps são tipos de referência
- Ordem dos elementos não é garantida
- Valor zero de um map é nil
- Devem ser inicializados antes do uso (make() ou literal)

OPERAÇÕES COMUNS:
- Inserção: mapa[chave] = valor
- Busca: valor, existe := mapa[chave]
- Remoção: delete(mapa, chave)
- Iteração: for chave, valor := range mapa
*/

func main() {
	fmt.Printf("=== DEMONSTRAÇÃO DE ARRAYS E MAPS EM GO ===\n")

	// ========== ARRAYS ==========
	fmt.Println("1. ARRAYS:")
	exemploArrays()

	fmt.Println("\n" + strings.Repeat("=", 50) + "\n")

	// ========== SLICES ==========
	fmt.Println("2. SLICES:")
	exemploSlices()

	fmt.Println("\n" + strings.Repeat("=", 50) + "\n")

	// ========== MAPS ==========
	fmt.Println("3. MAPS:")
	exemploMaps()

	fmt.Println("\n" + strings.Repeat("=", 50) + "\n")

	// ========== EXEMPLOS PRÁTICOS ==========
	fmt.Println("4. EXEMPLOS PRÁTICOS:")
	exemplosPraticos()
}

func exemploArrays() {
	// Declaração e inicialização de arrays
	fmt.Println("--- Declaração e Inicialização ---")

	// Array com tamanho fixo
	var numeros [5]int
	fmt.Printf("Array vazio: %v\n", numeros)

	// Array inicializado com valores
	idades := [3]int{25, 30, 35}
	fmt.Printf("Array inicializado: %v\n", idades)

	// Array com tamanho inferido
	nomes := [...]string{"Ana", "Bruno", "Carlos"}
	fmt.Printf("Array com tamanho inferido: %v\n", nomes)

	// Acessando elementos
	fmt.Println("\n--- Acessando Elementos ---")
	fmt.Printf("Primeiro nome: %s\n", nomes[0])
	fmt.Printf("Última idade: %d\n", idades[len(idades)-1])

	// Modificando elementos
	numeros[0] = 10
	numeros[4] = 50
	fmt.Printf("Array modificado: %v\n", numeros)

	// Iterando sobre arrays
	fmt.Println("\n--- Iteração ---")
	for i, nome := range nomes {
		fmt.Printf("Índice %d: %s\n", i, nome)
	}
}

func exemploSlices() {
	// Declaração e inicialização de slices
	fmt.Println("--- Declaração e Inicialização ---")

	// Slice vazio
	var frutas []string
	fmt.Printf("Slice vazio: %v (len: %d, cap: %d)\n", frutas, len(frutas), cap(frutas))

	// Slice inicializado
	cores := []string{"vermelho", "azul", "verde"}
	fmt.Printf("Slice inicializado: %v (len: %d, cap: %d)\n", cores, len(cores), cap(cores))

	// Slice com make
	numeros := make([]int, 3, 5) // tamanho 3, capacidade 5
	fmt.Printf("Slice com make: %v (len: %d, cap: %d)\n", numeros, len(numeros), cap(numeros))

	// Adicionando elementos
	fmt.Println("\n--- Adicionando Elementos ---")
	frutas = append(frutas, "maçã", "banana", "laranja")
	fmt.Printf("Após append: %v\n", frutas)

	// Slice de slice (sub-slice)
	fmt.Println("\n--- Sub-slices ---")
	subCores := cores[1:3] // do índice 1 até 2 (3 não incluído)
	fmt.Printf("Sub-slice cores[1:3]: %v\n", subCores)

	// Copiando slices
	fmt.Println("\n--- Copiando Slices ---")
	copia := make([]string, len(frutas))
	copy(copia, frutas)
	fmt.Printf("Cópia: %v\n", copia)

	// Removendo elementos (simulação)
	fmt.Println("\n--- Removendo Elementos ---")
	// Remove o elemento no índice 1
	frutas = append(frutas[:1], frutas[2:]...)
	fmt.Printf("Após remoção: %v\n", frutas)
}

func exemploMaps() {
	// Declaração e inicialização de maps
	fmt.Println("--- Declaração e Inicialização ---")

	// Map vazio
	var idades map[string]int
	fmt.Printf("Map vazio (nil): %v\n", idades)

	// Inicializando com make
	idades = make(map[string]int)
	fmt.Printf("Map inicializado: %v\n", idades)

	// Map literal
	capitais := map[string]string{
		"Brasil":    "Brasília",
		"França":    "Paris",
		"Japão":     "Tóquio",
		"Argentina": "Buenos Aires",
	}
	fmt.Printf("Map literal: %v\n", capitais)

	// Adicionando elementos
	fmt.Println("\n--- Adicionando Elementos ---")
	idades["Eduardo"] = 34
	idades["Bento"] = 20
	idades["Maria"] = 28
	fmt.Printf("Idades: %v\n", idades)

	// Acessando elementos
	fmt.Println("\n--- Acessando Elementos ---")
	idade := idades["Eduardo"]
	fmt.Printf("Idade do Eduardo: %d\n", idade)

	// Verificando se chave existe
	if idade, existe := idades["João"]; existe {
		fmt.Printf("Idade do João: %d\n", idade)
	} else {
		fmt.Println("João não encontrado no map")
	}

	// Iterando sobre maps
	fmt.Println("\n--- Iteração ---")
	for pais, capital := range capitais {
		fmt.Printf("%s: %s\n", pais, capital)
	}

	// Removendo elementos
	fmt.Println("\n--- Removendo Elementos ---")
	delete(idades, "Bento")
	fmt.Printf("Após remover Bento: %v\n", idades)

	// Map com tipos complexos
	fmt.Println("\n--- Map com Tipos Complexos ---")
	estudantes := map[string][]int{
		"Ana":    {85, 90, 78},
		"Bruno":  {92, 88, 95},
		"Carlos": {76, 82, 89},
	}
	fmt.Printf("Notas dos estudantes: %v\n", estudantes)
}

func exemplosPraticos() {
	fmt.Println("--- Sistema de Inventário ---")
	inventario := make(map[string]int)

	// Adicionando produtos
	produtos := []string{"notebook", "mouse", "teclado", "monitor"}
	quantidades := []int{10, 25, 15, 8}

	for i, produto := range produtos {
		inventario[produto] = quantidades[i]
	}

	fmt.Println("Inventário inicial:")
	for produto, qtd := range inventario {
		fmt.Printf("- %s: %d unidades\n", produto, qtd)
	}

	// Simulando vendas
	fmt.Println("\nApós vendas:")
	inventario["notebook"] -= 3
	inventario["mouse"] -= 10

	for produto, qtd := range inventario {
		status := "OK"
		if qtd < 10 {
			status = "ESTOQUE BAIXO"
		}
		fmt.Printf("- %s: %d unidades [%s]\n", produto, qtd, status)
	}

	fmt.Println("\n--- Análise de Frequência de Palavras ---")
	texto := []string{"go", "é", "uma", "linguagem", "go", "é", "rápida", "go", "é", "simples"}
	frequencia := make(map[string]int)

	// Contando frequência
	for _, palavra := range texto {
		frequencia[palavra]++
	}

	// Criando slice para ordenação
	type palavraFreq struct {
		palavra string
		freq    int
	}

	var palavras []palavraFreq
	for p, f := range frequencia {
		palavras = append(palavras, palavraFreq{p, f})
	}

	// Ordenando por frequência (decrescente)
	sort.Slice(palavras, func(i, j int) bool {
		return palavras[i].freq > palavras[j].freq
	})

	fmt.Println("Frequência de palavras (ordenada):")
	for _, pf := range palavras {
		fmt.Printf("'%s': %d vezes\n", pf.palavra, pf.freq)
	}

	fmt.Println("\n--- Agrupamento de Dados ---")
	pessoas := []struct {
		nome   string
		idade  int
		cidade string
	}{
		{"Ana", 25, "São Paulo"},
		{"Bruno", 30, "Rio de Janeiro"},
		{"Carlos", 25, "São Paulo"},
		{"Diana", 30, "São Paulo"},
		{"Eduardo", 25, "Rio de Janeiro"},
	}

	// Agrupando por idade
	porIdade := make(map[int][]string)
	for _, pessoa := range pessoas {
		porIdade[pessoa.idade] = append(porIdade[pessoa.idade], pessoa.nome)
	}

	fmt.Println("Pessoas agrupadas por idade:")
	for idade, nomes := range porIdade {
		fmt.Printf("Idade %d: %v\n", idade, nomes)
	}

	// Agrupando por cidade
	porCidade := make(map[string][]string)
	for _, pessoa := range pessoas {
		porCidade[pessoa.cidade] = append(porCidade[pessoa.cidade], pessoa.nome)
	}

	fmt.Println("\nPessoas agrupadas por cidade:")
	for cidade, nomes := range porCidade {
		fmt.Printf("%s: %v\n", cidade, nomes)
	}
}

/*
=== CONCEITOS IMPORTANTES ADICIONAIS ===

DIFERENÇAS ENTRE ARRAYS E SLICES:
1. Arrays têm tamanho fixo, slices são dinâmicos
2. Arrays são valores (copiados quando passados), slices são referências
3. Arrays raramente são usados diretamente, slices são mais comuns
4. Slices têm capacidade (cap) além do tamanho (len)

PERFORMANCE:
- Arrays: Acesso O(1), inserção/remoção não aplicável
- Slices: Acesso O(1), append O(1) amortizado, inserção/remoção O(n)
- Maps: Acesso O(1) médio, inserção/remoção O(1) médio

CASOS DE USO:
Arrays:
- Quando o tamanho é conhecido e fixo
- Para estruturas matemáticas (matrizes, vetores)
- Quando performance de memória é crítica

Slices:
- Listas dinâmicas
- Quando o tamanho varia
- Processamento de dados sequenciais

Maps:
- Índices/lookups rápidos
- Contadores e frequências
- Agrupamento de dados
- Cache simples

ARMADILHAS COMUNS:
1. Slice de slice compartilha array subjacente
2. Maps não são thread-safe
3. Iteração em maps não tem ordem garantida
4. Slice nil vs slice vazio são diferentes
5. Capacidade de slice pode crescer exponencialmente

BOAS PRÁTICAS:
1. Use slices ao invés de arrays na maioria dos casos
2. Inicialize maps com make() ou literal
3. Verifique se chave existe antes de usar valor do map
4. Use cap() para pré-alocar slices quando souber o tamanho aproximado
5. Cuidado com vazamentos de memória em sub-slices
*/
