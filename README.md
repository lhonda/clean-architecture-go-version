[![Go Report Card](https://goreportcard.com/badge/github.com/lhonda/clean-architecture-go-version)](https://goreportcard.com/report/github.com/lhonda/clean-architecture-go-version)

# clean-architecture-go-version

## Objetivo:

Este repositório é baseado no https://github.com/eminetto/clean-architecture-go-v2.

Criar uma POC contemplando os itens abaixo:

- Entidades:

- [x] Order(customer, pizza)
- [x] Pizza: lista de ingredientes


- Use case:

- [x] listar pizzas
- [x] criar pedido(Order)

- Controller
- [x] API
- [x] CLI

- Framework/Drivers

- Cobra (https://github.com/spf13/cobra)
- Fiber (https://github.com/gofiber/fiber)

Apresentação: API, CLI.

## Dependências

Para teste da API usamos o https://httpie.io.

Para atualizar as dependências do Go basta executar `go mod tidy`.

## Execução:

Para rodar o docker compose:

`chmod 666 ./ops/db/init.sql`

`docker-compose up`


## Comandos

### CLI

- listar pizzas:

`go run main.go list-pizzas`

- criar pedido:

`go run main.go create-order --owner Guido --pizzas presuto,queijo`

### API

No diretório api, executar a seguinte comando para rodar o servidor:

`go run main.go`

- listar pizzas

`http http://127.0.0.1:3000/pizzas`

Saída:
```HTTP/1.1 200 OK
Content-Length: 355
Content-Type: application/json
Date: Mon, 27 Sep 2021 18:10:46 GMT

[
    {
        "CreatedAt": "2021-09-27T17:53:55Z",
        "ID": "4c7966aa-33dd-4e5a-8038-46bd34d21e9f",
        "Ingredients": [
            "queijo;molho"
        ],
        "Name": "queijo",
        "Order": "4c7966aa-33dd-4e5a-8038-46bd34d21e9d"
    },
    {
        "CreatedAt": "2021-09-27T17:53:55Z",
        "ID": "4c7966aa-33dd-3e5a-8038-46bd34d21e9f",
        "Ingredients": [
            "presunto;molho"
        ],
        "Name": "presunto",
        "Order": "4c7966aa-33dd-4e5a-8038-46bd34d21e92"
    }
]
```

- criar pedido

`http http://127.0.0.1:3000/orders < order.json `

Saída:

```HTTP/1.1 201 Created
Content-Length: 461
Content-Type: application/json
Date: Mon, 27 Sep 2021 18:14:04 GMT

{
    "CreatedAt": "0001-01-01T00:00:00Z",
    "ID": "49af66b1-d263-47d1-aaa4-b61d84972496",
    "Owner": "Guido",
    "Pizzas": [
        {
            "CreatedAt": "2021-09-14T19:02:49Z",
            "ID": "4c7966aa-33dd-4e5a-8038-46bd34d21e9f",
            "Ingredients": [
                "queijo;molho"
            ],
            "Name": "queijo",
            "Order": "4c7966aa-33dd-4e5a-8038-46bd34d21e9d"
        },
        {
            "CreatedAt": "2021-09-14T19:02:49Z",
            "ID": "4c7966aa-33dd-3e5a-8038-46bd34d21e9f",
            "Ingredients": [
                "presunto;molho"
            ],
            "Name": "presunto",
            "Order": "4c7966aa-33dd-4e5a-8038-46bd34d21e92"
        }
    ]
}
```

## Testes

Gerar dados do code coverage:

`go test -coverprofile=.coverage.out ./...`

Gera relatório de code coverage:

`go tool cover -html=.coverage.out`

## Comentários

Nesta sessão listamos os tópicos interessantes do uso de Clean Architecture nos nossos projetos.

1.Existe a possibilidade de isolar o domínio em um projeto; deste modo a camada de API/cmd pode importar e usar o domínio.

2.Podemos adicionar o Swagger/Open API, gitlab/github badges.

3.Constatamos que modelar sistemas em Go é trabalhoso no início do projeto.

4.A lib Testify ajuda na leitura dos testes.

5.Fizemos reuniões semanais de uma hora. Começamos pelas entidades, use cases, persistência e camadas externas. Todas as camadas tem testes de unidade.

6.A codificação foi feita em grupo, houve troca de experiência, resolução de bugs, discussão da arquitetura. A discussão em grupo e o compartilhamento de experiência foi importante pois acreditamos que pair programming aumenta a qualidade do código e difunde a cultura interna do projeto entre os membros do time.



## Referências:

- https://github.com/eminetto/clean-architecture-go-v2.

- https://www.youtube.com/watch?v=M0ODYSahNq8
