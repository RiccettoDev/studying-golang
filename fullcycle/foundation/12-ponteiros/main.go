package main

import "fmt"

// ============================================
// CONCEITO: Ponteiros em Go
// ============================================
// Ponteiro é uma variável que armazena o ENDEREÇO DE MEMÓRIA de outra variável
// Operadores:
//   & (ampersand) - retorna o endereço de memória de uma variável
//   * (asterisco)  - desreferencia um ponteiro (acessa o valor)

// ============================================
// Exemplo 1: Básico - Valor vs Ponteiro
// ============================================

func exemploBasico() {
	fmt.Println("=== Exemplo 1: Básico ===")

	// Variável normal
	x := 10
	fmt.Printf("Valor de x: %d\n", x)
	fmt.Printf("Endereço de x: %p\n", &x) // & retorna o endereço

	// Ponteiro para x
	var p *int // declara ponteiro para int
	p = &x     // p recebe o endereço de x
	fmt.Printf("Valor do ponteiro p: %p\n", p)
	fmt.Printf("Valor apontado por p: %d\n", *p) // * desreferencia

	// Modificando através do ponteiro
	*p = 20
	fmt.Printf("Novo valor de x: %d\n", x) // x foi modificado!
}

// ============================================
// Exemplo 2: Passagem por Valor vs Ponteiro
// ============================================

func modificarPorValor(n int) {
	n = 100 // modifica apenas a cópia
	fmt.Println("Dentro da função (valor):", n)
}

func modificarPorPonteiro(n *int) {
	*n = 100 // modifica o valor original
	fmt.Println("Dentro da função (ponteiro):", *n)
}

func exemploPassagemParametro() {
	fmt.Println("\n=== Exemplo 2: Passagem de Parâmetros ===")

	// Por valor
	a := 10
	fmt.Println("Antes (valor):", a)
	modificarPorValor(a)
	fmt.Println("Depois (valor):", a) // não mudou!

	// Por ponteiro
	b := 10
	fmt.Println("\nAntes (ponteiro):", b)
	modificarPorPonteiro(&b)
	fmt.Println("Depois (ponteiro):", b) // mudou!
}

// ============================================
// Exemplo 3: Ponteiros com Structs
// ============================================

type Pessoa struct {
	Nome  string
	Idade int
}

// Método com receiver por valor (recebe cópia)
func (p Pessoa) AniversarioPorValor() {
	p.Idade++
	fmt.Printf("Dentro do método (valor): %s tem %d anos\n", p.Nome, p.Idade)
}

// Método com receiver por ponteiro (modifica original)
func (p *Pessoa) AniversarioPorPonteiro() {
	p.Idade++
	fmt.Printf("Dentro do método (ponteiro): %s tem %d anos\n", p.Nome, p.Idade)
}

func exemploStructs() {
	fmt.Println("\n=== Exemplo 3: Ponteiros com Structs ===")

	// Por valor
	p1 := Pessoa{Nome: "João", Idade: 30}
	fmt.Printf("Antes: %s tem %d anos\n", p1.Nome, p1.Idade)
	p1.AniversarioPorValor()
	fmt.Printf("Depois: %s tem %d anos\n", p1.Nome, p1.Idade) // não mudou!

	// Por ponteiro
	p2 := Pessoa{Nome: "Maria", Idade: 25}
	fmt.Printf("\nAntes: %s tem %d anos\n", p2.Nome, p2.Idade)
	p2.AniversarioPorPonteiro()                               // Go automaticamente converte para &p2
	fmt.Printf("Depois: %s tem %d anos\n", p2.Nome, p2.Idade) // mudou!
}

// ============================================
// Exemplo 4: new() - Alocação de Memória
// ============================================

func exemploNew() {
	fmt.Println("\n=== Exemplo 4: Função new() ===")

	// new() aloca memória e retorna um ponteiro
	p := new(int)
	fmt.Printf("Ponteiro p: %p\n", p)
	fmt.Printf("Valor inicial: %d\n", *p) // zero value de int = 0

	*p = 42
	fmt.Printf("Novo valor: %d\n", *p)

	// Com struct
	pessoa := new(Pessoa)
	fmt.Printf("Pessoa vazia: %+v\n", pessoa)
	pessoa.Nome = "Carlos"
	pessoa.Idade = 35
	fmt.Printf("Pessoa preenchida: %+v\n", pessoa)
}

// ============================================
// Exemplo 5: Ponteiro Nil
// ============================================

func exemploNil() {
	fmt.Println("\n=== Exemplo 5: Ponteiro Nil ===")

	var p *int
	fmt.Printf("Ponteiro não inicializado: %v\n", p) // nil

	if p == nil {
		fmt.Println("Ponteiro é nil!")
	}

	// ⚠️ Desreferenciar ponteiro nil causa panic
	// fmt.Println(*p) // PANIC: invalid memory address

	// Inicializando
	x := 10
	p = &x
	fmt.Printf("Ponteiro inicializado: %v, valor: %d\n", p, *p)
}

// ============================================
// Exemplo 6: Slice e Map (já são referências)
// ============================================

func modificarSlice(s []int) {
	s[0] = 999 // modifica o slice original!
}

func modificarMap(m map[string]int) {
	m["chave"] = 100 // modifica o map original!
}

func exemploReferencias() {
	fmt.Println("\n=== Exemplo 6: Slice e Map (referências) ===")

	// Slice
	numeros := []int{1, 2, 3}
	fmt.Println("Antes:", numeros)
	modificarSlice(numeros)
	fmt.Println("Depois:", numeros) // mudou! slice é referência

	// Map
	dados := map[string]int{"chave": 10}
	fmt.Println("\nAntes:", dados)
	modificarMap(dados)
	fmt.Println("Depois:", dados) // mudou! map é referência
}

// ============================================
// Exemplo 7: Ponteiro para Ponteiro
// ============================================

func exemploPonteiroParaPonteiro() {
	fmt.Println("\n=== Exemplo 7: Ponteiro para Ponteiro ===")

	x := 10
	p := &x  // ponteiro para x
	pp := &p // ponteiro para ponteiro

	fmt.Printf("Valor de x: %d\n", x)
	fmt.Printf("Endereço de x: %p\n", &x)
	fmt.Printf("Valor de p (endereço de x): %p\n", p)
	fmt.Printf("Valor apontado por p: %d\n", *p)
	fmt.Printf("Endereço de p: %p\n", &p)
	fmt.Printf("Valor de pp (endereço de p): %p\n", pp)
	fmt.Printf("Valor apontado por pp: %p\n", *pp)
	fmt.Printf("Valor apontado por *pp: %d\n", **pp)

	// Modificando através do ponteiro para ponteiro
	**pp = 50
	fmt.Printf("Novo valor de x: %d\n", x)
}

// ============================================
// Exemplo 8: Quando usar Ponteiro vs Valor
// ============================================

type PequenoStruct struct {
	ID int
}

type GrandeStruct struct {
	Dados [1000]int
	Nome  string
	Info  map[string]string
}

func processarPorValor(g GrandeStruct) {
	// Copia toda a struct (pesado!)
	_ = g
}

func processarPorPonteiro(g *GrandeStruct) {
	// Passa apenas o endereço (leve!)
	_ = g
}

func exemploPerformance() {
	fmt.Println("\n=== Exemplo 8: Performance ===")

	grande := GrandeStruct{
		Nome: "Dados grandes",
		Info: make(map[string]string),
	}

	fmt.Println("Struct grande criada")
	fmt.Println("Por valor: copia toda a struct (lento)")
	processarPorValor(grande)

	fmt.Println("Por ponteiro: passa apenas endereço (rápido)")
	processarPorPonteiro(&grande)
}

// ============================================
// Exemplo 9: Retornando Ponteiros
// ============================================

func criarPessoa(nome string, idade int) *Pessoa {
	p := Pessoa{Nome: nome, Idade: idade}
	return &p // ✅ Seguro! Go move p para heap se necessário
}

func exemploRetornarPonteiro() {
	fmt.Println("\n=== Exemplo 9: Retornando Ponteiros ===")

	pessoa := criarPessoa("Ana", 28)
	fmt.Printf("Pessoa criada: %+v\n", pessoa)
	fmt.Printf("Endereço: %p\n", pessoa)
}

// ============================================
// Exemplo 10: Comparação de Ponteiros
// ============================================

func exemploComparacao() {
	fmt.Println("\n=== Exemplo 10: Comparação ===")

	x := 10
	y := 10

	p1 := &x
	p2 := &x
	p3 := &y

	fmt.Printf("p1 == p2: %v (mesmo endereço)\n", p1 == p2)
	fmt.Printf("p1 == p3: %v (endereços diferentes)\n", p1 == p3)
	fmt.Printf("*p1 == *p3: %v (valores iguais)\n", *p1 == *p3)
}

func main() {
	exemploBasico()
	exemploPassagemParametro()
	exemploStructs()
	exemploNew()
	exemploNil()
	exemploReferencias()
	exemploPonteiroParaPonteiro()
	exemploPerformance()
	exemploRetornarPonteiro()
	exemploComparacao()

	fmt.Println("\n=== Resumo ===")
	fmt.Println("& - retorna endereço de memória")
	fmt.Println("* - desreferencia ponteiro (acessa valor)")
	fmt.Println("Use ponteiros para: modificar valores, evitar cópias, trabalhar com structs grandes")
}
