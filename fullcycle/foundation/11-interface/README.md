# Interfaces em Go

## 🎯 O que é uma Interface?

Uma interface em Go é um **contrato** que define um conjunto de métodos. Qualquer tipo que implemente esses métodos automaticamente satisfaz a interface.

## 🔑 Características principais

### 1. Implementação Implícita

Diferente de Java/C#, em Go você **não declara** que um tipo implementa uma interface:

```go
// ❌ Não existe isso em Go
type Client implements Pessoa { ... }

// ✅ Se Client tem os métodos, ele automaticamente implementa
type Pessoa interface {
    Desativar()
}

type Client struct { ... }

func (c *Client) Desativar() { ... }  // Client agora implementa Pessoa!
```

### 2. Duck Typing

> "Se anda como um pato e faz quack como um pato, então é um pato"

```go
type Animal interface {
    Falar() string
}

// Qualquer tipo com método Falar() string é um Animal
```

## 📚 Tipos de Interfaces

### Interface Vazia (`interface{}` ou `any`)

Aceita **qualquer tipo**:

```go
func Print(i interface{}) {
    fmt.Println(i)
}

Print("string")
Print(42)
Print(true)
Print(struct{}{})
```

A partir do Go 1.18, você pode usar `any` (alias para `interface{}`):

```go
func Print(i any) {
    fmt.Println(i)
}
```

### Interface com Métodos

Define comportamentos específicos:

```go
type Writer interface {
    Write([]byte) (int, error)
}

type Reader interface {
    Read([]byte) (int, error)
}

// Composição de interfaces
type ReadWriter interface {
    Reader
    Writer
}
```

## 🔍 Type Assertion

Extrair o valor concreto de uma interface:

### Forma segura (com verificação)

```go
func ProcessarValor(i interface{}) {
    if v, ok := i.(string); ok {
        fmt.Println("É uma string:", v)
    } else {
        fmt.Println("Não é uma string")
    }
}
```

### Forma direta (pode causar panic)

```go
func ProcessarValor(i interface{}) {
    v := i.(string)  // ⚠️ Panic se não for string
    fmt.Println(v)
}
```

## 🔀 Type Switch

Verificar múltiplos tipos:

```go
func IdentificarTipo(i interface{}) {
    switch v := i.(type) {
    case string:
        fmt.Println("String:", v)
    case int:
        fmt.Println("Int:", v)
    case bool:
        fmt.Println("Bool:", v)
    default:
        fmt.Printf("Tipo desconhecido: %T\n", v)
    }
}
```

## 🎭 Polimorfismo

Interfaces permitem polimorfismo em Go:

```go
type Forma interface {
    Area() float64
}

type Retangulo struct {
    Largura, Altura float64
}

type Circulo struct {
    Raio float64
}

func (r Retangulo) Area() float64 {
    return r.Largura * r.Altura
}

func (c Circulo) Area() float64 {
    return 3.14 * c.Raio * c.Raio
}

func ImprimirArea(f Forma) {
    fmt.Println("Área:", f.Area())
}

func main() {
    r := Retangulo{Largura: 10, Altura: 5}
    c := Circulo{Raio: 7}

    ImprimirArea(r)  // Funciona!
    ImprimirArea(c)  // Funciona!
}
```

## ⚠️ Ponteiro vs Valor

**Importante:** Se o método usa receiver por ponteiro, você precisa passar ponteiro:

```go
type Pessoa interface {
    Desativar()
}

type Client struct {
    Active bool
}

// Método com receiver por ponteiro
func (c *Client) Desativar() {
    c.Active = false
}

func main() {
    c := Client{Active: true}

    // ❌ Erro: Client não implementa Pessoa
    // var p Pessoa = c

    // ✅ Correto: *Client implementa Pessoa
    var p Pessoa = &c
    p.Desativar()
}
```

## 🌟 Interfaces Comuns da Standard Library

### io.Reader

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}
```

Implementado por: `os.File`, `strings.Reader`, `bytes.Buffer`, etc.

### io.Writer

```go
type Writer interface {
    Write(p []byte) (n int, err error)
}
```

Implementado por: `os.File`, `bytes.Buffer`, `http.ResponseWriter`, etc.

### fmt.Stringer

```go
type Stringer interface {
    String() string
}
```

Usado por `fmt.Println` para converter tipos em string:

```go
type Pessoa struct {
    Nome string
    Idade int
}

func (p Pessoa) String() string {
    return fmt.Sprintf("%s (%d anos)", p.Nome, p.Idade)
}

func main() {
    p := Pessoa{Nome: "João", Idade: 30}
    fmt.Println(p)  // Saída: João (30 anos)
}
```

### error

```go
type error interface {
    Error() string
}
```

Qualquer tipo com método `Error() string` é um erro:

```go
type MeuErro struct {
    Mensagem string
}

func (e MeuErro) Error() string {
    return e.Mensagem
}
```

## 💡 Boas Práticas

### 1. Interfaces pequenas

> "The bigger the interface, the weaker the abstraction" - Rob Pike

```go
// ✅ Bom - interface pequena e focada
type Reader interface {
    Read([]byte) (int, error)
}

// ❌ Ruim - interface muito grande
type SuperInterface interface {
    Read() error
    Write() error
    Close() error
    Open() error
    Delete() error
    Update() error
}
```

### 2. Aceite interfaces, retorne structs

```go
// ✅ Bom
func ProcessarDados(r io.Reader) *Resultado {
    // ...
}

// ❌ Ruim
func ProcessarDados(f *os.File) io.Reader {
    // ...
}
```

### 3. Defina interfaces onde são usadas

```go
// ✅ Bom - interface definida no package que usa
package meuapp

type Storage interface {
    Save(data []byte) error
}

func ProcessarArquivo(s Storage) { ... }

// ❌ Ruim - interface definida no package que implementa
package database

type Storage interface {
    Save(data []byte) error
}
```

## 🎓 Resumo

| Conceito                    | Descrição                                            |
| --------------------------- | ---------------------------------------------------- |
| **Implementação implícita** | Não precisa declarar `implements`                    |
| **Interface vazia**         | `interface{}` ou `any` aceita qualquer tipo          |
| **Type assertion**          | Extrair valor concreto: `v, ok := i.(string)`        |
| **Type switch**             | Verificar múltiplos tipos com `switch v := i.(type)` |
| **Polimorfismo**            | Diferentes tipos implementando mesma interface       |
| **Ponteiro vs valor**       | Atenção ao receiver do método                        |
| **Interfaces pequenas**     | Preferir interfaces focadas e específicas            |

## 🚀 Vantagens das Interfaces em Go

1. **Desacoplamento** - código menos dependente de implementações concretas
2. **Testabilidade** - fácil criar mocks e stubs
3. **Flexibilidade** - adicionar novos tipos sem modificar código existente
4. **Composição** - combinar interfaces pequenas em maiores
5. **Simplicidade** - implementação implícita reduz boilerplate
