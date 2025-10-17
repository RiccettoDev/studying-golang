# Organização de pastas e arquivos em um projeto

- Todo projeto se iniia com o comando **go mod init**, assim como temos o **package.json** em node, em golang temos **go.mod**

```
go mod init github.com/riccettodev/studying-golang
```

- Usamos este link de diretório do github para nomear o pacote do projeto

- Está é uma boa prática para criação do projeto.

## Pasta Internal

- Tudo que é inerente a sua aplicação, tudo que é apenas de uso interno de sua aplicação, deve ficar dentro desta pasta.

- Nome das pastas e dos arquivos sempre no singular
- Sempre com letras minúsculas
- E sempre separados por underline

-> no arquivo de teste, se mantiver o mesmo nome do pacote, terá acesso a todas as variáveis, métodos e funções privadas. Para seu teste, caso não querira ter acesso é preciso alterar o nome do pacote, por exemplo mudar de **package entity** para **package entity_test**

- apesar de tudo no singular as pastas docs, configs e scripts vão no plural

- O que se guarda na pasta api, são os arquivos de geração da api, como por exemplo o Swagger Documentation | Swagger Docs.

- Na pasta test da raiz, é onde ficam testes e2e, pois os unitários ficam ao lado do arquivo que vai ser testado mesmo.

## Exemplo de estrutura:

```
.
├── api
├── build
│   ├── docker-compose.yaml
│   └── Dockerfile
├── cmd
│   ├── api-client
│   │   └── main.go
│   └── api-server
│       └── main.go
├── configs
├── docs
│   └── aula20.md
├── go.mod
├── internal
│   └── domain
│       └── entity
│           ├── entity_test.go
│           └── entity.go
├── scripts
└── test
```
