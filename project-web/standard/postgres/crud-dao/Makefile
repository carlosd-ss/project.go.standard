# Makefile
.EXPORT_ALL_VARIABLES:	

GO111MODULE=on
GOPROXY=direct
GOSUMDB=off
PORT_SERVER = 5209

build:
	@echo "########## Compilando nossa API ... "
	#go build -ldflags="-s -w" -o crud-dao.go
	CGO_ENABLED=0 GOOS=linux GOARCH=386 go build --trimpath -a -tags netgo -ldflags '-w -extldflags "-static"' -o crud.user.singleton main.go
	@echo "buid completo..."
	@echo "\033[0;33m################ run #####################\033[0m"
	./crud.user.singleton

update:
	@echo "########## Compilando nossa API ... "
	@rm -f go.*
	go mod init github.com/jeffotoni/project.go.standard/project-web/standard/postgres/crud-dao
	CGO_ENABLED=0 GOOS=linux GOARCH=386 go build --trimpath -a -tags netgo -ldflags '-w -extldflags "-static"' -o crud.user.singleton main.go
	@echo "buid completo..."
	@echo "fim"

