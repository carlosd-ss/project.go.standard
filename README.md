
# Project Go Standart

![GO](https://img.shields.io/badge/GO-%2300ADD8?style=for-the-badge&logo=Go&labelColor=%23444444)
![Postgres
](https://img.shields.io/badge/POSTGRESQL-%23336791?style=for-the-badge&logo=PostgreSQL&logoColor=%23336791&labelColor=%23444444) 
![MONGODB](https://img.shields.io/badge/MONGO-%237A248?style=for-the-badge&logo=MongoDB&labelColor=%23444444)

Este repositório irá ajuda-lo a decidir qual melhor layout a ser utilizado em um projeto na linguagem Go.

O que estamos tentando fazer por aqui é documentar e apresentar alguns modelos mais utilizados que percebemos ao longo dos anos usando a linguagem Go. **Não** temos a pretenção de determinar o que é melhor ou pior mas temos como apresenta-lo alguns exemplos práticos e suas abordagens de como aplica-los e somente você saberá o que é melhor para o seu projeto.

Todo este repositório foi feito baseado na linguagem Go, da nossa humilde experiência de utilização em nosso dia a dia como Linguagem principal em nossos projetos, então você está convidado a nos enviar issues, pull requests o que achar necessário para que possamos melhorar nosso repositório de layouts padrões para projetos em Go. Quando iniciamos em Go e precisamos fazer algo muito simples, nem é necessário um padrão ou layout isto torna-se um exagero eu diria desnesessário talvez um main.go já resolveria seu problema, sempre da preferência para o mais simples, o mais enxuto possível é um caminho interessante a seguir. 

Quando seu projeto envolver mais pessoas para colaborar, uma equipe etc, e o projeto começa a crescer neste cenário iremos precisar de um padrão arquitetural ou um layout para organizar nossos projetos em Go. O que irá encontrar neste repositório são diversos modelos de layout e organização de projetos em diversos cenários e aplicabilidades. Não temos a pretenção de apresentar todas as possibilidades isto seria insano, mas vamos mostrar alguns que poderam ajuda-lo e talvez até poderá criar o seu próprio modelo e padrão para seus projetos.



## 🕮 Alguns diretorios usados nos layouts e suas estruturas

### Projetos Web Standard ☑️

```_bash
├── Makefile
├── Dockerfile
├── README.md
├── go.mod
├── go.sum
├── main.go
├── certs
├── config
├── controler
│    ├── handler
│    │   ├── user
│    │       └── post.user.go
│    │       └── post.user_test.go
│    │       └── put.user.go
│    │       └── put.user_test.go
│    │       └── delete.user.go
│    │       └── delete.user_test.go
│    │       └── get.user.go
│    │       └── get.user_test.go
│    ├── middleware
│         └── logger.go
│         └── gzip.go
│         └── cors.go
├── route
│     ├── user
│         └── user.go
├── model
│     ├── user
│         └── user.go
├── pkg
│   └── fmts
│       └── fmts.go
│       └── fmts_test.go
└── repo
    ├── user
        └── add.user.go
        └── add.user_test.go
        └── del.user.go
        └── del.user_test.go
        └── up.user.go
        └── up.user_test.go
        └── get.user.go
        └── get.user_test.go
```  
### Projetos Web Clean Arquitecture ☑️

```_bash
├── Makefile
├── README.md
├── app
│   ├── domain
│   │   ├── model
│   │   ├── repository
│   │   └── service
│   ├── interface
│   │   ├── persistence
│   │   └── rpc
│   ├── registry
│   └── usecase
├── cmd
│   └── 8am
│       └── main.go
└── vendor
    ├── vendor packages
```    

### Projetos Web microservice Grpc ☑️

```_bash
├── Makefile
├── README.md
├── certs
├── cmd
│   ├── cli
│   ├── server
│   ├── gw
├── docker
│   └── Dockerfile
│   └── Makefile
├── pkg
│   └── fmts
│   └── grpc
└── proto
    ├── user
    │     └── user.go
    │     └── user.proto
    ├── customer
    │     └── customer.go
    │     └── customer.proto
```    

### Projetos Web fragment service 


```_bash
├── Makefile
├── README.md
├── config
├── model
│    ├── user
│         └── user.go
├── pkg
│   └── fmts
│       └── fmts.go
│       └── fmts_test.go
```   

### Projetos Serverless


```_bash
├── Makefile
├── README.md
├── model
│    ├── user
│         └── user.go
├── pkg
│   └── fmts
│       └── fmts.go
│       └── fmts_test.go
```   



### Projetos Lib 


```_bash
├── Makefile
├── README.md
```   



## 🔗 Links relacionados


[📘 Go Documentação](https://golang.org/doc/)

[📘 Go Faq](https://golang.org/doc/faq)

[📘 Go Tour](https://tour.golang.org/welcome/1)

[📘 Go Efetivo](https://golang.org/doc/effective_go.html)