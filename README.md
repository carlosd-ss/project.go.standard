
![Go](https://github.com/jeffotoni/project.go.standard/blob/master/.github/go.png)


# Project Go Standard
![GO](https://img.shields.io/badge/GO-%2300ADD8?style=for-the-badge&logo=Go&labelColor=%23444444)
![Postgres
](https://img.shields.io/badge/POSTGRESQL-%23336791?style=for-the-badge&logo=PostgreSQL&logoColor=%23336791&labelColor=%23444444) 
![MONGODB](https://img.shields.io/badge/MONGO-%237A248?style=for-the-badge&logo=MongoDB&labelColor=%23444444)

Este repositório irá ajuda-lo a decidir um melhor layout a ser adotado em seus projeto utilizandos na linguagem Go 😍.

O que estamos tentando fazer por aqui é documentar e apresentar alguns modelos mais utilizados que percebemos ao longo dos anos usando a linguagem Go. **Não** temos a pretenção de determinar o que é melhor ou pior mas temos como apresenta-lo alguns exemplos práticos e suas abordagens de como aplica-los e somente você saberá o que é melhor para o seu projeto.

Todo este repositório foi feito baseado na linguagem Go 😍, da nossa humilde experiência de utilização em nosso dia a dia como Linguagem principal em nossos projetos, então você está convidado a nos enviar issues, pull requests o que achar necessário para que possamos melhorar nosso repositório de layouts padrões para projetos em Go. 

Quando iniciamos em Go e precisamos fazer algo muito simples, nem é necessário um padrão ou layout isto torna-se um exagero eu diria desnesessário talvez um main.go já resolveria seu problema, sempre da preferência para o mais simples, o mais enxuto possível a utilização de menas libs externas possíveis e inúmeras outras boas práticas que podem adotar em seus projetos sempre é um caminho interessante a seguir.

É interessante em um segundo momento é claro neste mesmo projeto, abordar algumas boas práticas
que podemos adotar utilizando Go como por exemplo utilizações e construções desnecessárias e mirabolates
um bom exemplo seria a utilização sem necessidade e de forma desacerbada quando o assunto é utilizar os ponteiros, mas a princípio iremos iniciar com
a organização do layout do seu projeto para depois avançarmos ainda mais e proporcionar
algo mais neste projeto apesar de que os exemplos práticos apresentados já estão de uma forma bem bacaninha mas longe da perfeição e caso percebam que pode melhorar por favor
nos envie issues ou pull requests ❤️.

Neste projeto existe um diretório específico para abordar os Patterns que podemos desenvolver utilizando Go é uma introdução que vamos enrriquecer ao longo do tempo com os exemplos práticos.

Voltando ao Layout e organização de projetos quando seu projeto envolver mais pessoas para colaborar, uma equipe etc, e o projeto começa a crescer neste cenário iremos precisar de um padrão arquitetural ou um layout para organizar nossos projetos em Go e um dos pontos mais importantes mantê-lo. O que irá encontrar neste repositório são diversos modelos de layout e organização de projetos em diversos cenários e aplicabilidades.

Não temos a pretenção de apresentar todas as possibilidades isto seria insano,
mas vamos mostrar alguns que poderam facilitar seu dia a dia em seus projetos
e talvez possamos até influencia-lo a criar o seu próprio modelo de layout e
tenha certeza que iriamos ficar super felizes com isto ❤️ e é claro compartilhe sempre conosco para colaborar com a comunidade.

## Alguns diretórios usados nos layouts e suas estruturas
****************************

### Projetos Web Standard Singleton ☑️

```_bash
├── config
├── controler
│    ├── handler
│    │   ├── user
│    │       └── ping.go
│    │       └── ping_test.go
│    │       └── get.param.go
│    │       └── user.post.go
│    │       └── user.post_test.go
│    │       └── user.put.go
│    │       └── user.put_test.go
│    │       └── user.delete.go
│    │       └── user.delete_test.go
│    │       └── user.get.go
│    │       └── user.get_test.go
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
├── pkg
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
├── pkg
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

### Projetos Web Standard DAO ☑️

```_bash
├── config
├── controler
│    ├── handler
│    │    └── ping.go
│    │    └── ping_test.go
│    │    └── model.connect.go
│    │    └── get.param.go
│    │    └── user.go
│    │    └── user_test.go
│    ├── middleware
│         └── mw.go
│  └── all.route.go
│  └── endpoints.go
│  └── model.server.go
│  └── show.route.go
│  └── stopserver.go
├── model
│     ├── user
│         └── user.go
├── pkg
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
        └── user.go
        └── user_test.go
├── Makefile
├── Dockerfile
├── README.md
├── go.mod
├── go.sum
├── main.go

```

### Projetos Web Standard DAO Clean ☑️

```_bash
├── config
├── controler
│    ├── handler
│    │    └── ping.go
│    │    └── ping_test.go
│    │    └── model.connect.go
│    │    └── get.param.go
│    │    └── user.go
│    │    └── user_test.go
│    ├── middleware
│         └── mw.go
│  └── route.go
├── model
│     ├── user
│         └── user.go
├── pkg
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
        └── user.go
        └── user_test.go
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

### Projetos Web microservice Grpc ☑️

```_bash
├── certs
│     └── server.crt
│     └── server.csr
│     └── server.key
├── cmd
│   ├── cli         
│   │    └── cli.go
│   ├── server
│   │        └── server.go 
│   ├── gw
│       └── gw.go    
├── docker
│   └── Dockerfile
│   └── Makefile
├── pkg
│   └── fmts
│   └── grpc
└── proto
│   ├── user
│   │     └── user.go
│   │     └── user.proto
│   ├── customer
│   │     └── customer.go
│   │     └── customer.proto
├── Makefile
├── Dockerfile
├── README.md
├── go.mod
├── go.sum
├── main.go
└──
```    

### Projetos Web fragment service 

```_bash
├── certs
├── config
├── controler
│    ├── handler
│    │     └──  user
│    │           └── post.user.go
│    │           └── post.user_test.go
│    │       
│    ├── middleware
│         └── logger.go
│         └── gzip.go
│         └── cors.go
├── route
│     └──  user
│         └── user.go
├── model
│     └──  user
│         └── user.go
├── pkg
│   └── fmts
│       └── fmts.go
│       └── fmts_test.go
└── repo
     └── user
        └── add.user.go
        └── add.user_test.go
├── Makefile
├── Dockerfile
├── README.md
├── go.mod
├── go.sum
├── main.go
└──

```  


### Projetos Serverless


```_bash
├── Makefile
├── README.md
├── Project.User
│   ├── go.mod
│   ├── go.sum
│   ├── main.go
│   ├── Makefile
│   ├── models
│   │   ├── user.go
│   └── pkg
│       ├── user
│       │   └── publicar.go
│       └── sqs
│           └── sendjobfifo.go
├── Project.Customer
│   ├── go.mod
│   ├── go.sum
│   ├── main.go
│   ├── Makefile
│   ├── models
│   │   ├── sqs.customer.go
│   │   ├── return.customer.go
│   └── pkg
│       ├── customer
│       │   └── publicar.go
│       └── sqs
│           └── sendjoberror.go
└──
```   



### Projetos Lib 


```_bash
├── Makefile
├── README.md
├── Dockerfile
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
