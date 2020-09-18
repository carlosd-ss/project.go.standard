# Makefile
.EXPORT_ALL_VARIABLES:	

GO111MODULE=on
GOPROXY=direct
GOSUMDB=off
PORT_SERVER = 5209

build:
	@echo "########## Compilando nossa API ... "
	go build -ldflags="-s -w" -o api0.go
	@echo "buid completo..."
	@echo "\033[0;33m################ run #####################\033[0m"

update:
	@echo "########## Compilando nossa API ... "
	@rm -f go.*
	go mod init github.com/jeffotoni/project.go.standard/project.web/standard.libray/crud.user.singleton
	go build -trimpath -ldflags="-s -w" -o api0.go
	sleep 1
	@echo "buid completo..."
	@echo "fim"

