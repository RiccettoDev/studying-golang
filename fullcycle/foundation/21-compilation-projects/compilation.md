# Compilação de Projetos em Go

Go é uma linguagem compilada que gera um binário nativo — sem precisar de runtime ou VM para executar.

## Comandos principais

- `go build` — compila o projeto e gera um executável no diretório atual.
- `go build -o nome_do_binario` — define o nome do arquivo de saída.
- `go run main.go` — compila e executa sem gerar arquivo permanente.
- `go install` — compila e coloca o binário em `$GOPATH/bin`.

## Cross-compilation (compilar para outro S.O.)

Go facilita compilar para qualquer plataforma usando as variáveis `GOOS` e `GOARCH`:

```bash
GOOS=linux GOARCH=amd64 go build -o app-linux main.go
GOOS=windows GOARCH=amd64 go build -o app.exe main.go
GOOS=darwin GOARCH=arm64 go build -o app-mac main.go
```

Valores comuns: `GOOS` = linux, windows, darwin | `GOARCH` = amd64, arm64, 386.

## Dicas úteis

- O binário gerado é **estático** por padrão — não depende de libs externas no sistema alvo.
- Use `go build -ldflags="-s -w"` para reduzir o tamanho do binário (remove debug info).
- Liste todas as plataformas suportadas com: `go tool dist list`.
- Para projetos com múltiplos arquivos, rode `go build ./...` na raiz do módulo.
