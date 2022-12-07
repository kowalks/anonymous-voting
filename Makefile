.PHONY: build build-cs build-all p1 p2 p3

CMD:=./cmd
BIN:=./bin

MANAGER:=manager

build:
	go build -o $(BIN)/$(MANAGER) $(CMD)/$(MANAGER)/main.go

build-all: build

manager:
	go run $(CMD)/$(MANAGER)/main.go