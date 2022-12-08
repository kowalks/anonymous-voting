package main

import (
	"fmt"
	"math/big"
	"os"

	"github.com/kowalks/anonymous-voting/pkg/contract"
	"github.com/kowalks/anonymous-voting/pkg/math"
)

func requestInfo(name string, rnd, rGx, rGy *big.Int) (*big.Int, *big.Int) {
	fmt.Println("Key Manager", name)
	fmt.Printf("    r = %v", rnd)
	fmt.Println()
	fmt.Println("   rG =", rGx, rGy)

	fmt.Println()
	fmt.Print("Announced: ")

	xStr, yStr := "0", "0"
	fmt.Scan(&xStr, &yStr)
	x := new(big.Int)
	x, _ = x.SetString(xStr, 10)
	y := new(big.Int)
	y, _ = y.SetString(yStr, 10)

	fmt.Println()

	return x, y
}

func revealSecret(instance *contract.Contract, address string, contractAddress string, r *big.Int) {
	phase, _ := instance.Phase(nil)
	for phase != 2 {
		phase, _ = instance.Phase(nil)
	}

	auth, _ := contract.Auth(address)
	instance = contract.Connect(contractAddress)
	instance.RevealPublicKey(auth, r)

	fmt.Println("Secret revealed!")
	fmt.Println()
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Provide a name and a rand int for this key manager")
		os.Exit(0)
	}

	// Generating random number for Key Manager
	name := os.Args[1]
	rnd := math.GenerateRand(os.Args[2])

	// Connecting to blockchain
	auth, instance, address, cAddress := contract.ConnectionCLI()

	// Generating pubkey and registering manager
	rGx, rGy := math.EllipticCurve.ScalarBaseMult(rnd.Bytes())
	tx, err := instance.RegisterManager(auth, rGx, rGy)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Registered manager with tx", tx.Hash())
	fmt.Println()

	math.PrintCurve(math.EllipticCurve)
	x, y := requestInfo(name, rnd, rGx, rGy)
	rrGx, rrGy := math.EllipticCurve.ScalarMult(x, y, rnd.Bytes())
	fmt.Println("Public Key rrG =", rrGx, rrGy)

	// Announcing first pubkey for the election
	auth, _ = contract.Auth(address)
	instance = contract.Connect(cAddress)
	_, err = instance.AnnouncePublicKey(auth, rrGx, rrGy)
	if err != nil {
		panic(err)
	}

	fmt.Println("Annouced public key")
	fmt.Println()

	// Revealing pubkey after votation
	revealSecret(instance, address, cAddress, rnd)
}
