 # Crud 

| Banco    | Formato |
|----------|---------|
| Postgres | DAO     |

```bash
├── controller
│   ├── handler
│   │   ├── customer
│   │   │   ├── delete.go
│   │   │   ├── get.all.go
│   │   │   ├── get.uuid.go
│   │   │   ├── post.go
│   │   │   ├── server.go
│   │   │   └── update.go
│   │   └── ping
│   │       └── ping.go
│   ├── middleware
│   │   └── logger
│   │       └── config.go 
│   └── route
│       └── route.go
├── go.mod
├── go.sum
├── main.go
├── Makefile
├── models
│   ├── customer
│   │   └── customer.go
│   ├── errors
│   │   └── erros.go
│   └── ping
│       └── ping.go
├── pkg
│   ├── fmts
│   │   ├── fmts.go
│   │   └── fmts_test.go
│   ├── logone
│   │   └── logone.go
│   ├── psql
│   │   ├── config.go
│   │   └── connect.go
│   ├── regex
│   │   └── regex.go
│   └── validador
│       └── validador.go
├
└── repo
    └── customer
        ├── delete.go
        ├── get.all.go
        ├── get.uuid.go
        ├── post.go
        └── update.go


 ```
