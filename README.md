
![Go](https://github.com/carlosd-ss/project.go.standard/blob/master/.github/go.png)


# Project Go Standard
![GO](https://img.shields.io/badge/GO-%2300ADD8?style=for-the-badge&logo=Go&labelColor=%23444444)
![Postgres
](https://img.shields.io/badge/POSTGRESQL-%23336791?style=for-the-badge&logo=PostgreSQL&logoColor=%23336791&labelColor=%23444444) 
![MONGODB](https://img.shields.io/badge/MONGO-%237A248?style=for-the-badge&logo=MongoDB&labelColor=%23444444)

Este repositório irá ajuda-lo a decidir um melhor layout a ser adotado em seus projeto utilizandos na linguagem Go 😍.

O que estamos tentando fazer por aqui é documentar e apresentar alguns modelos mais utilizados que percebemos ao longo dos anos usando a linguagem Go. **Não** temos a pretenção de determinar o que é melhor ou pior mas temos como apresenta-lo alguns exemplos práticos e suas abordagens de como aplica-los e somente você saberá o que é melhor para o seu projeto.

## Alguns diretórios usados nos layouts e suas estruturas
****************************

### Projetos Web Standard DAO ☑️

```_bash
├── config
├── controler
│    ├── handler
│    │    └── ping.go
│    │    └── ping_test.go
│    │    └── model.connect.go
│    │    └── get.param.go
│    │    └── user.post.go
│    │    └── user.post_test.go
│    │    └── user.put.go
│    │    └── user.put_test.go
│    │    └── user.delete.go
│    │    └── user.delete_test.go
│    │    └── user.get.go
│    │    └── user.get_test.go
│    ├── middleware
│         └── adapter.go
│         └── basic.go
│         └── cors.go
│         └── custom-header.go
│         └── gzip.go
│         └── maxclient.go
│         └── gjwt.go
│  └── all.route.go
│  └── endpoints.go
│  └── model.server.go
│  └── show.route.go
│  └── stopserver.go
├── model
│     ├── user
│         └── user.go
├── internal
│   └── cert
│   └── cors
│   └── crypt
│   └── gjwt
│   └── psql
│   └── util
│   └── zerolog
│   └── fmts
│       └── fmts.go
│       └── fmts_test.go
└── postman
└── pgdmp
└── repo
    ├── user
        └── user.add.go
        └── user.add_test.go
        └── user.del.go
        └── user.del_test.go
        └── user.up.go
        └── user.up_test.go
        └── user.get.go
        └── user.get_test.go
├── Makefile
├── Dockerfile
├── README.md
├── go.mod
├── go.sum
├── main.go

```


### Projetos Web Clean Arquitecture ☑️

```_bash
├── app
│   ├── domain
│   │   ├── model            
│   │   │     └── user.go
│   │   ├── router
│   │   │      └── user.go 
│   │   ├── mocks
│   │   │     └── user.go 
│   │   │     └── user_test.go 
│   │   ├── repository
│   │   │     └── cert
│   │   │     └── cors
│   │   │     └── crypt
│   │   │     └── gjwt
│   │   │     └── psql
│   │   │     └── util
│   │   │     └── zerolog
│   │   │     └── fmts
│   │   │          └── fmts.go
│   │   │          └── fmts_test.go
│   │   └── service
│   │           └── service.go
│   │           └── service_test.go
│   │
│   ├── interface
│   │   ├── persistence
│   │   └── rpc
│   ├── registry
│   └── usecase
├── cmd
│    └── main.go
└── vendor
│   ├── packages
├── Makefile
├── Dockerfile
├── README.md
├── go.mod
├── go.sum
├── main.go
└──
```       

## 🔗 Links relacionados


[📘 Go Documentação](https://golang.org/doc/)

[📘 Go Faq](https://golang.org/doc/faq)

[📘 Go Tour](https://tour.golang.org/welcome/1)

[📘 Go Efetivo](https://golang.org/doc/effective_go.html)
