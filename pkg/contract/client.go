package contract

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var hostURL = "http://localhost:8545"

func Deploy(clientAddress string) *Contract {
	auth, client := Auth(clientAddress)

	address, tx, instance, err := DeployContract(auth, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Contract deployed!")
	fmt.Println("  contract address:", address.Hex())
	fmt.Println("  transaction hash:", tx.Hash().Hex())

	return instance
}

func Connect(contractAddress string) *Contract {
	client, err := ethclient.Dial(hostURL)
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress(contractAddress)
	instance, err := NewContract(address, client)
	if err != nil {
		log.Fatal(err)
	}

	return instance
}

func Auth(clientAddress string) (*bind.TransactOpts, *ethclient.Client) {
	client, err := ethclient.Dial(hostURL)
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA(clientAddress)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatal(err)
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = gasPrice

	return auth, client
}

func ConnectionCLI() (*bind.TransactOpts, *Contract, string, string) {
	fmt.Println("Client setup: provide client address and contract address")

	fmt.Print("  address: ")
	var address string
	fmt.Scan(&address)

	fmt.Print("  contract: ")
	var contract string
	fmt.Scan(&contract)

	fmt.Println()

	auth, _ := Auth(address)
	instance := Connect(contract)

	return auth, instance, address, contract
}
