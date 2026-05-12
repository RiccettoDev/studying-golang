# Ponteiros em Go

## 🎯 O que é um Ponteiro?

Um ponteiro é uma variável que armazena o **endereço de memória** de outra variável, em vez de armazenar o valor diretamente.

```
Memória:
┌─────────┬─────────┐
│ 0x1000  │   10    │  ← variável x
├─────────┼─────────┤
│ 0x2000  │ 0x1000  │  ← ponteiro p (guarda endereço de x)
└─────────┴─────────┘
```

## 🔧 Operadores

| Operador | Nome        | Função                        | Exemplo       |
| -------- | ----------- | ----------------------------- | ------------- |
| `&`      | Address-of  | Retorna o endereço de memória | `p := &x`     |
| `*`      | Dereference | Acessa o valor apontado       | `valor := *p` |

## 📝 Sintaxe Básica

```go
// Declaração
var p *int        // ponteiro para int (valor inicial: nil)

// Atribuição
x := 10
p = &x            // p recebe o endereço de x

// Acesso ao valor
fmt.Println(*p)   // imprime 10

// Modificação
*p = 20           // x agora vale 20
```

## 🔄 Passagem por Valor vs Ponteiro

### Por Valor (cópia)

```go
func modificar(n int) {
    n = 100  // modifica apenas a cópia
}

x := 10
modificar(x)
fmt.Println(x)  // ainda é 10
```

### Por Ponteiro (referência)

```go
func modificar(n *int) {
    *n = 100  // modifica o original
}

x := 10
modificar(&x)
fmt.Println(x)  // agora é 100
```

## 🏗️ Ponteiros com Structs

### Receiver por Valor

```go
type Pessoa struct {
    Nome  string
    Idade int
}

func (p Pessoa) Aniversario() {
    p.Idade++  // modifica apenas a cópia
}

pessoa := Pessoa{Nome: "João", Idade: 30}
pessoa.Aniversario()
fmt.Println(pessoa.Idade)  // ainda é 30
```

### Receiver por Ponteiro

```go
func (p *Pessoa) Aniversario() {
    p.Idade++  // modifica o original
}

pessoa := Pessoa{Nome: "João", Idade: 30}
pessoa.Aniversario()  // Go converte automaticamente para &pessoa
fmt.Println(pessoa.Idade)  // agora é 31
```

## 🆕 Alocação com `new()`

A função `new()` aloca memória e retorna um ponteiro:

```go
// Com new()
p := new(int)
*p = 42

// Equivalente a:
var x int
p := &x
*p = 42
```

Para structs:

```go
// Com new()
pessoa := new(Pessoa)
pessoa.Nome = "Maria"

// Equivalente a:
p := Pessoa{}
pessoa := &p
pessoa.Nome = "Maria"

// Ou mais comum:
pessoa := &Pessoa{Nome: "Maria"}
```

## ⚠️ Ponteiro Nil

Um ponteiro não inicializado tem valor `nil`:

```go
var p *int
fmt.Println(p)  // <nil>

if p == nil {
    fmt.Println("Ponteiro não inicializado")
}

// ❌ PANIC! Não desreferencie ponteiro nil
// fmt.Println(*p)  // runtime error: invalid memory address
```

## 🔗 Tipos de Referência

Alguns tipos em Go já são **referências** por natureza:

| Tipo      | Comportamento                  |
| --------- | ------------------------------ |
| `slice`   | Referência (modifica original) |
| `map`     | Referência (modifica original) |
| `channel` | Referência (modifica original) |
| `array`   | Valor (cria cópia)             |
| `struct`  | Valor (cria cópia)             |

```go
func modificarSlice(s []int) {
    s[0] = 999  // modifica o slice original!
}

numeros := []int{1, 2, 3}
modificarSlice(numeros)
fmt.Println(numeros)  // [999 2 3]
```

## 🎭 Ponteiro para Ponteiro

É possível ter ponteiros que apontam para outros ponteiros:

```go
x := 10
p := &x      // *int
pp := &p     // **int

fmt.Println(x)    // 10
fmt.Println(*p)   // 10
fmt.Println(**pp) // 10

**pp = 50
fmt.Println(x)    // 50
```

## 🚀 Quando Usar Ponteiros?

### ✅ Use Ponteiros quando:

1. **Precisa modificar o valor original**

```go
func Incrementar(n *int) {
    *n++
}
```

2. **Struct grande (evitar cópia)**

```go
type GrandeStruct struct {
    Dados [10000]int
}

func Processar(g *GrandeStruct) {  // ✅ Eficiente
    // ...
}
```

3. **Métodos que modificam o receiver**

```go
func (p *Pessoa) Envelhecer() {
    p.Idade++
}
```

4. **Precisa representar "ausência de valor"**

```go
var nome *string  // nil = sem nome
```

### ❌ Não use Ponteiros quando:

1. **Tipos pequenos (int, bool, string pequena)**

```go
func Somar(a, b int) int {  // ✅ Por valor é mais simples
    return a + b
}
```

2. **Slice, map, channel (já são referências)**

```go
func ProcessarSlice(s []int) {  // ✅ Não precisa de *[]int
    // ...
}
```

3. **Métodos que só leem dados**

```go
func (p Pessoa) GetNome() string {  // ✅ Por valor é suficiente
    return p.Nome
}
```

## 🔍 Comparação de Ponteiros

```go
x := 10
y := 10

p1 := &x
p2 := &x
p3 := &y

fmt.Println(p1 == p2)   // true (mesmo endereço)
fmt.Println(p1 == p3)   // false (endereços diferentes)
fmt.Println(*p1 == *p3) // true (valores iguais)
```

## 🧠 Escape Analysis

Go decide automaticamente se uma variável vai para **stack** ou **heap**:

```go
func criarPessoa() *Pessoa {
    p := Pessoa{Nome: "João"}
    return &p  // ✅ Go move p para heap automaticamente
}
```

Não se preocupe com isso - o compilador otimiza!

## 📊 Tabela Resumo

| Conceito          | Sintaxe           | Descrição                   |
| ----------------- | ----------------- | --------------------------- |
| Declarar ponteiro | `var p *int`      | Ponteiro para int           |
| Obter endereço    | `p = &x`          | p recebe endereço de x      |
| Acessar valor     | `*p`              | Valor apontado por p        |
| Modificar valor   | `*p = 10`         | Altera valor apontado       |
| Alocar memória    | `p := new(int)`   | Cria int e retorna ponteiro |
| Ponteiro nil      | `var p *int`      | Ponteiro não inicializado   |
| Receiver ponteiro | `func (p *T) M()` | Método modifica original    |
| Receiver valor    | `func (p T) M()`  | Método recebe cópia         |

## 💡 Dicas Importantes

1. **Go não tem aritmética de ponteiros** (diferente de C/C++)

```go
p := &x
p++  // ❌ Erro de compilação
```

2. **Sintaxe simplificada para structs**

```go
p := &Pessoa{Nome: "João"}
p.Nome = "Maria"  // não precisa (*p).Nome
```

3. **Conversão automática em métodos**

```go
pessoa := Pessoa{}
pessoa.Aniversario()  // Go converte para &pessoa automaticamente
```

4. **Zero value de ponteiro é nil**

```go
var p *int
fmt.Println(p == nil)  // true
```

## 🎯 Regra de Ouro

> "Use ponteiros quando precisar modificar ou quando a cópia for cara. Caso contrário, use valores."

## 🔗 Relação com Outros Conceitos

- **Interfaces**: Se o método usa receiver por ponteiro, só `*T` implementa a interface
- **Goroutines**: Cuidado com race conditions ao compartilhar ponteiros
- **Garbage Collector**: Go gerencia memória automaticamente, não precisa `free()`

## ✅ Checklist de Uso

- [ ] Preciso modificar o valor original? → Use ponteiro
- [ ] Struct é grande (>100 bytes)? → Use ponteiro
- [ ] É slice/map/channel? → Não use ponteiro
- [ ] Método só lê dados? → Use valor
- [ ] Preciso representar "nulo"? → Use ponteiro (nil)
