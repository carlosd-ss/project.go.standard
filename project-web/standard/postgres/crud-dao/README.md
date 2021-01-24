 # Crud 

| Banco    | Formato |
|----------|---------|
| Postgres | DAO     |

```bash

├── config
│   └── config.go
├── controller
│   ├── all.route.go
│   ├── endpoint.go
│   ├── handler
│   │   ├── ping.go
│   │   ├── user.delete.go
│   │   ├── user.get.go
│   │   ├── user.param.go
│   │   ├── user.post.go
│   │   └── user.put.go
│   ├── middleware
│   │   ├── mw-adapter.go
│   │   ├── mw-basic.go
│   │   ├── mw-cors.go
│   │   ├── mw-custom-header.go
│   │   ├── mw.gjwt.go
│   │   ├── mw-gzip.go
│   │   ├── mw-logger.go
│   │   └── mw-maxclient.go
│   ├── model.server.go
│   ├── show.route.go
│   └── stopserver.go
├── Dockerfile
├── go.mod
├── go.sum
├── internal
│   ├── cert
│   │   └── cert.go
│   ├── cors
│   │   ├── cors.go
│   │   └── utils.go
│   ├── crypt
│   │   └── crypt.go
│   ├── fmts
│   │   ├── fmts.go
│   │   └── fmts_test.go
│   ├── gjwt
│   │   ├── config.go
│   │   └── gjwt.go
│   ├── psql
│   │   └── connect.go
│   ├── util
│   │   ├── file.go
│   │   ├── regex-email.go
│   │   ├── regex-endpoint.go
│   │   ├── regex-phone.go
│   │   └── string.go
│   └── zerolog
│       └── zerolog.go
├── main.go
├── Makefile
├── model
│   └── user
│       └── user.go
├── pgdmp
│   └── gostandard.sql
├── postman
│   └── go.standard.lib.postman_collection.json
├
└── repo
    └── user
        ├── user.delete.go
        ├── user.get.go
        ├── user.post.go
        └── user.put.go

 ```
