package main

import (
	"crypto/rand"
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
	return x, y
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Provide a name and a rand int for this key manager")
		os.Exit(0)
	}

	// Generating random number for Key Manager
	name := os.Args[1]
	rnd := new(big.Int)
	rnd, ok := rnd.SetString(os.Args[2], 10)
	if !ok {
		panic("Error in converting number to big.Int")
	}

	if len(rnd.Bytes()) == 0 {
		order := math.EllipticCurve.Params().N
		rnd, _ = rand.Int(rand.Reader, order)
	}

	// Connecting to blockchain
	auth, client, instance := contract.ConnectionCLI()

	// Generating pubkey
	rGx, rGy := math.EllipticCurve.ScalarBaseMult(rnd.Bytes())

	// Announcing pubkey
	tx, _ := instance.AnnouncePublicKey(auth, rGx, rGy)
	fmt.Println("Announced pubkey with tx", tx.Hash())

	math.PrintCurve(math.EllipticCurve)
	x, y := requestInfo(name, rnd, rGx, rGy)

	rrGx, rrGy := math.EllipticCurve.ScalarMult(x, y, rnd.Bytes())
	fmt.Println("  rrG =", rrGx, rrGy)

	_ = client
}
