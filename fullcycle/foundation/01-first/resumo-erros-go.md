# Go — Conceitos aplicados no módulo `01-first`

Este módulo demonstra conceitos fundamentais do Go através de exemplos práticos e erros comuns.

## 📁 Estrutura do projeto

```
01-first/
├── go.mod            → module 01-first
├── hello/
│   └── hello.go      → package hello (subpackage)
├── main.go           → package main (ponto de entrada)
└── ola.go            → package main (mesmo package do main)
```

## 🎯 Conceitos aplicados

### 1️⃣ **Múltiplos arquivos no mesmo package**

**ola.go**

```go
package main

const b = "Olá Mundo!"
```

**main.go**

```go
package main

import (
	"01-first/hello"
	"fmt"
)

func main() {
	println(b)              // ✅ Acessa constante de ola.go
	fmt.Println(hello.A)    // ✅ Acessa constante do package hello
}
```

**📌 Conceito:** Todos os arquivos `.go` na mesma pasta com `package main` fazem parte do mesmo programa. A constante `b` definida em `ola.go` está disponível em `main.go` automaticamente.

### 2️⃣ **Criação de subpackages**

**hello/hello.go**

```go
package hello

const A = "Hello Golang!"
```

**📌 Conceito:** Criamos um subpackage `hello` dentro de uma subpasta. Para ser acessível de outros packages, a constante usa letra maiúscula (`A`).

### 3️⃣ **Visibilidade (exported vs unexported)**

| Identificador   | Visibilidade                      | Exemplo           |
| --------------- | --------------------------------- | ----------------- |
| `b` (minúscula) | ❌ Privado (só dentro do package) | `const b = "..."` |
| `A` (maiúscula) | ✅ Público (exportado)            | `const A = "..."` |

**No código:**

- `b` em `ola.go` → acessível apenas dentro do `package main`
- `A` em `hello.go` → acessível de qualquer package via `hello.A`

### 4️⃣ **Sistema de módulos (go.mod)**

**go.mod**

```
module 01-first

go 1.25.2
```

**📌 Conceito:** O `go.mod` define o nome do módulo (`01-first`). Isso permite importar subpackages usando o caminho completo:

```go
import "01-first/hello"
```

### 5️⃣ **Diferença entre `println` e `fmt.Println`**

```go
println(b)              // função built-in (não precisa import)
fmt.Println(hello.A)    // função do package fmt (precisa import)
```

**📌 Conceito:**

- `println` → função nativa do Go, útil para debug rápido
- `fmt.Println` → função do package padrão, mais robusta e formatada

## ❌ Erros comuns e soluções

### Erro 1: `expected 'package', found 'EOF'`

**Causa:** Arquivo vazio ou sem declaração de package.

**Solução:** Todo arquivo `.go` precisa começar com:

```go
package nome
```

### Erro 2: `undefined: b`

**Causa:** Executar apenas um arquivo do package:

```bash
go run main.go  # ❌ Ignora ola.go
```

**Solução:** Executar o package completo:

```bash
cd 01-first
go run .        # ✅ Compila todos os arquivos do package main
```

### Erro 3: `cannot find package "01-first/hello"`

**Causa:** Falta o arquivo `go.mod` ou está com nome errado.

**Solução:** Inicializar módulo:

```bash
go mod init 01-first
```

## 🚀 Como executar

```bash
cd fullcycle/foundation/01-first
go run .
```

**Saída esperada:**

```
Olá Mundo!
Hello Golang!
```

## ✅ Resumo dos conceitos

| Conceito                                | Aplicação no código                              |
| --------------------------------------- | ------------------------------------------------ |
| **Múltiplos arquivos no mesmo package** | `main.go` + `ola.go` compartilham `package main` |
| **Subpackages**                         | `hello/hello.go` é um package separado           |
| **Visibilidade**                        | `b` (privado) vs `A` (exportado)                 |
| **Módulos**                             | `go.mod` define `module 01-first`                |
| **Imports**                             | `import "01-first/hello"` para usar subpackage   |
| **Execução correta**                    | `go run .` compila todos os arquivos do package  |
