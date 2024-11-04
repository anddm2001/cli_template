.PHONY: config build-lin build-win build-mac
.DEFAULT_GOAL := build-lin

PATH_SOURCE=./main.go
PATH_BIN=./cli_example
PATH_BIN_WIN=./cli_example.exe 

config:
	rm -rf .env
	cp .env.example .env

build-lin:
	CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w" -o $(PATH_BIN) $(PATH_SOURCE)

build-win:
	CGO_ENABLED=0 GOOS=windows go build -ldflags "-s -w" -o $(PATH_BIN_WIN) $(PATH_SOURCE)

build-mac:
	CGO_ENABLED=0 GOOS=darwin go build -ldflags "-s -w" -o $(PATH_BIN) $(PATH_SOURCE)