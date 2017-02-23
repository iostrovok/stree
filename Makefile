# vim: set softtabstop=2 shiftwidth=2:
SHELL = bash

export GOPATH := ${PWD}
CURDIR = ${PWD}

all: install start

install:
	@echo "Build" 
	$(shell rm ./bin/main)
	$(shell export GOPATH=${CURDIR}; go build -o ./bin/main ./src/main.go)

start:
	@echo "Start" 
	./bin/main

test:
	@echo ${GOPATH}
	go get gopkg.in/check.v1
	@go test ./src/symmetrictree
