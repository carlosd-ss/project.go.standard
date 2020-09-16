#!/bin/bash
# Go Api server
# @jeffotoni
# 2019-06-01

echo "-------------------------------------- Clean <none> images ---------------------------------------"
docker rmi $(docker images | grep "<none>" | awk '{print $3}') --force

echo "\033[0;33m################################## go build crud.user.singleton.api.user ##################################\033[0m"
GOOS=linux go build -ldflags="-s -w" -o api.user main.go

echo "\033[0;33m################################## build docker crud.user.singleton.api.user ##################################\033[0m"
docker build -f Dockerfile -t 873761630739.dkr.ecr.us-east-1.amazonaws.com/crud.user.singleton.api.user .

echo "\033[0;33m######################################### login aws and push ########################################\033[0m"
$(aws ecr get-login --no-include-email --region us-east-1)
docker push 873761630739.dkr.ecr.us-east-1.amazonaws.com/crud.user.singleton.api.user
echo "\033[0;32mGenerated\033[0m \033[0;33m[ok]\033[0m \n"
