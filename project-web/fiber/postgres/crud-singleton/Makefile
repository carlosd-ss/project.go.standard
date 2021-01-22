# Makefile
.EXPORT_ALL_VARIABLES:

GO111MODULE=on
GOPROXY=direct
GOSUMDB=off

update:
	@echo "########## Compilando nossa API ... "
	rm -f go.*
	go mod init github.com/go.standard.project.layout/project.web/fiber/crud.postgresa
	go build --trimpath -ldflags="-s -w" -o crud.postgres main.go
	#upx crud.postgres
	@echo "buid completo..."
	@echo "\033[0;33m################ run #####################\033[0m"
	./crud.postgres

build:
	@echo "########## Compilando nossa API ... "
	go build --trimpath -ldflags="-s -w" -o crud.postgres main.go
	#upx crud.postgres
	@echo "buid completo..."
	@echo "\033[0;33m################ run #####################\033[0m"
	./crud.postgres