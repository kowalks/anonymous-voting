package main

import (
	"fmt"
	"os"

	"github.com/kowalks/anonymous-voting/pkg/contract"
	"github.com/kowalks/anonymous-voting/pkg/math"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Provide a name and a rand int for this candidate")
		os.Exit(0)
	}

	// Generating random number for candidate
	name := os.Args[1]
	rnd := math.GenerateRand(os.Args[2])

	// Connecting to blockchain
	auth, client, instance := contract.ConnectionCLI()

	// Generating pubkey and registering candidate
	x, y := math.EllipticCurve.ScalarBaseMult(rnd.Bytes())
	_, err := instance.AddCandidate(auth, name, x, y)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Get candidate count
	cnt, err := instance.CandidateCount(nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Candidate count:", cnt)

	_ = client
}
