package main

import (
	"fmt"
	"os"

	"github.com/kowalks/anonymous-voting/pkg/contract"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Provide an address to deploy the contract!")
		os.Exit(0)
	}

	// fixing address from with contract is deployed
	address := os.Args[1]

	instance := contract.Deploy(address)

	cnt, _ := instance.CandidateCount(nil)
	fmt.Println("Candidate count:", cnt)

	cnt, _ = instance.VoterCount(nil)
	fmt.Println("Voter count:", cnt)
}
