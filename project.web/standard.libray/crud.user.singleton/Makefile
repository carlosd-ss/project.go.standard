# Makefile
.EXPORT_ALL_VARIABLES:	

GO111MODULE=on
GOPROXY=direct
GOSUMDB=off
PORT_SERVER = 5209

build:
	@echo "########## Compilando nossa API ... "
	go build -ldflags="-s -w" -o api.user main.go
	@echo "buid completo..."
	@echo "\033[0;33m################ run #####################\033[0m"
	./api.user

deploy:
	@echo "########## Compilando nossa API ... "
	go build -ldflags="-s -w" -o api.user main.go
	@echo "buid completo..."
	@echo "\033[0;33m################ Enviando para o server #####################\033[0m"
	scp -P 2022 api.user root@apideploy.s3wf.com.br:/tmp
	rm -f api.user
	@echo "fim"

update:
	@echo "########## Compilando nossa API ... "
	@rm -f go.*
	go mod init github.com/jeffotoni/project.go.standard/project.web/standard.libray/crud.user/api.user
	go build -trimpath -ldflags="-s -w" -o api.user main.go
	sleep 1
	@echo "buid completo..."
	@echo "fim"

deploy.aws:
	@echo "########## Compilando nossa API ... "
	sh deploy-aws.sh
	@echo "fim"
