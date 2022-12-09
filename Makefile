.PHONY: build build-cs build-all p1 p2 p3

CMD:=./cmd
BIN:=./bin
FXT:=./fixtures

MANAGER:=manager
CANDIDATE:=candidate
SIGNER:=signer
VOTER:=voter
DEPLOY:=deployer

ADDRESS_REPLACER:='2s/.*/$(1)/'

build:
	go build -o $(BIN)/$(MANAGER) $(CMD)/$(MANAGER)/main.go
	go build -o $(BIN)/$(SIGNER) $(CMD)/$(SIGNER)/main.go
	go build -o $(BIN)/$(VOTER) $(CMD)/$(VOTER)/main.go

compile:
	solc --abi ./contracts/AnonymousVoting.sol -o bin --overwrite
	solc --bin ./contracts/AnonymousVoting.sol -o bin --overwrite
	abigen --bin=./bin/AnonymousVoting.bin --abi=./bin/AnonymousVoting.abi --pkg=contract --out=./pkg/contract/voting.go

address: $(FXT)/*
	cat bin/address.txt \
	 | awk -v q="'" '{print "sed -e "q"2s/.*/"$$1"/"q" -i "q q" $^"}' \
	 | bash

setup: compile setup-ganache address	

setup-ganache: teardown
	ganache-cli -m "ces27"  &
	sleep 2
	go run $(CMD)/$(DEPLOY).go cbed89197463cd2cfc71c2fe196989d844e56363269c11fc1b1fe510ed95f431 \
	 | grep contract \
	 | awk '{print $$3}' \
	 | cut -c3- > $(BIN)/address.txt

teardown:
	pgrep -f ganache-cli | xargs kill

candidates: c1 c2 c3

alice:
	go run $(CMD)/$(MANAGER).go Alice 1

bob:
	go run $(CMD)/$(MANAGER).go Bob 2

c1:
	go run $(CMD)/$(CANDIDATE).go Einstein 10 < $(FXT)/c1.txt

c2:
	go run $(CMD)/$(CANDIDATE).go Rutherford 11 < $(FXT)/c2.txt

c3:
	go run $(CMD)/$(CANDIDATE).go Thomson 12 < $(FXT)/c3.txt

voter:
	go run $(CMD)/$(VOTER).go
