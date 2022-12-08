.PHONY: build build-cs build-all p1 p2 p3

CMD:=./cmd
BIN:=./bin

MANAGER:=manager
SIGNER:=signer

build:
	go build -o $(BIN)/$(MANAGER) $(CMD)/$(MANAGER)/main.go
	go build -o $(BIN)/$(SIGNER) $(CMD)/$(SIGNER)/main.go

build-all: build

alice:
	go run $(CMD)/$(MANAGER)/main.go Alice 1

bob:
	go run $(CMD)/$(MANAGER)/main.go Bob 2

signer:
	go run $(CMD)/$(SIGNER)/main.go