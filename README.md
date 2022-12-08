# anonymous-voting

## Setting up environment

### Ganache

Install Ganache CLI tool for a fast, reliable local testnet.

```
npm install -g ganache-cli
```

### Solidity

Install Solidity compiler.

```
brew update
brew tap ethereum/ethereum
brew install solidity
```

Also install `abigen` tool

```
go get -u github.com/ethereum/go-ethereum
cd $GOPATH/src/github.com/ethereum/go-ethereum/
make
make devtools
```

### Running Ganache Testnet

The mnemonic "ces27" is added for deterministic account scheme generation.

```
ganache-cli -m "ces27"
```

The generated accounts should be the following

```
(0) 0xffae362fb776ce8b5A13D2b1ae679ECa0FD4A64e (100 ETH)
(1) 0x2164952E6Cd3214f67170F27209E8C22C4F70efc (100 ETH)
(2) 0x499f9d7Ba289B22cF71341a0B55cdb5eD929ac14 (100 ETH)
(3) 0xe199782E066b061c29097da5471707395126aF75 (100 ETH)
(4) 0x1eAD322628D33dc38660b15331AEBc5295D47278 (100 ETH)
(5) 0x1746EBe4a6FC1FB7081c9451779F3828B1b5F013 (100 ETH)
(6) 0x7726E00348b29Db7FDdC4e0100f143EAd01C1e76 (100 ETH)
(7) 0xdC27Ee49a65432e19CA21f240dac7C3A233375Ef (100 ETH)
(8) 0x7D4aFf221c308a00F1Ef15bb106125E8BfFe9cfE (100 ETH)
(9) 0x5eC1c5660BfF1012B061030E0bA5e3A07B124C62 (100 ETH)
```

The private keys are the following

```
(0) 0xcbed89197463cd2cfc71c2fe196989d844e56363269c11fc1b1fe510ed95f431
(1) 0x0ac0753be57e786048d3b2c09209485e290d88ba8d4a8cb5037dc79735f4fbd2
(2) 0xada477acd4db4d709a7ec2592c932ac3b19f948a93f468fb7c3c485a7a0b7b75
(3) 0x3705b1238edfa988aaecca5401df14abb11a7861e36b9996a80e57b0e3e301b4
(4) 0x3706b74cfeaf397c7c704c35f91711670e4124bd355b70d1b39679487ccbbb79
(5) 0x85f70d9a2f4dcb3cda5747fbec8555e7bb33cf47a67d195e46a5b0bbe4a368ab
(6) 0x33161f95f9c029d70e542d7bd13e1bcb677b8cb18436b95d0ac651841e68d0c5
(7) 0xa9a76319b6e5dbcf8afe2db55e76ab91a390efce633488982ecae9d625e3eb2a
(8) 0x7b06a0b3421c2447126477088e38cf32bc925717f74b3803b521ddc126b088b8
(9) 0x47b13371fb72c0f60f7dfa8c76fed2b39e8f2bdaefb9b3b42199251a2f898cad
```

### Compiling contracts

Compile the contracts.
```
solc --abi ./contracts/AnonymousVoting.sol -o bin
solc --bin ./contracts/AnonymousVoting.sol -o bin
abigen --bin=./bin/AnonymousVoting.bin --abi=./bin/AnonymousVoting.abi --pkg=eth --out=./pkg/eth/voting.go
```

## Deploying contracts

Now, we have to deploy our contract in the blockchain. For that, we have to specify an address with ETH to make the deployment. For simplicity, just choose `0xcbed89197463cd2cfc71c2fe196989d844e56363269c11fc1b1fe510ed95f431`

```
go run ./cmd/deployer.go cbed89197463cd2cfc71c2fe196989d844e56363269c11fc1b1fe510ed95f431
```

Everytime this script is executed, a new contract is deployed at a different address. If it's the first transaction in the blockchain, the addresses should be the following.
```
Connected to Ganache
Contract deployed!
  contract address: 0xd1285556883015FFFFe2C44227918E112290bdF0
  transaction hash: 0x9dd8f6d4eea464805ae49bb1baae78905545246ee22886a2598e488567642707
Candidate count: 0
Voter count: 0
```