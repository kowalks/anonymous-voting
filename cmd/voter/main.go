package main

import (
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"strconv"

	"github.com/kowalks/anonymous-voting/pkg/contract"
	"github.com/kowalks/anonymous-voting/pkg/math"
	"github.com/kowalks/anonymous-voting/pkg/sign"
)

func requestVote(instance *contract.Contract, address string, contractAddress string) (sign.Candidate, []sign.Candidate) {
	fmt.Println("Waiting for election to start")
	fmt.Println()

	phase, _ := instance.Phase(nil)
	for phase != 1 {
		phase, _ = instance.Phase(nil)
	}

	var candidates []sign.Candidate

	candidateCount, _ := instance.CandidateCount(nil)
	for i := big.NewInt(0); candidateCount.Cmp(i) == 1; i.Add(i, big.NewInt(1)) {
		candidate, _ := instance.Candidates(nil, i)
		name := candidate.Name
		x := candidate.Pubkey.X
		y := candidate.Pubkey.Y

		candidates = append(candidates, sign.Candidate{Name: name, Pubkey: &ecdsa.PublicKey{Curve: math.EllipticCurve, X: x, Y: y}})
	}

	fmt.Println("Candidate list")
	for i := 0; i < len(candidates); i++ {
		fmt.Printf("  (%d) %v\n", i, candidates[i].Name)
	}
	fmt.Println()
	fmt.Print("Choose an option: ")

	var option string
	fmt.Scan(&option)
	opt, err := strconv.Atoi(option)
	for err != nil || opt < 0 || opt >= len(candidates) {
		fmt.Print("Choose an option: ")
		fmt.Scan(&option)
		opt, err = strconv.Atoi(option)
	}

	return candidates[opt], candidates
}

func retriveParams(instance *contract.Contract, priv ecdsa.PrivateKey) ([]sign.Voter, []ecdsa.PublicKey, []ecdsa.PublicKey, int) {
	cnt, _ := instance.VoterCount(nil)
	curve := math.EllipticCurve

	var voters []sign.Voter
	var P, I []ecdsa.PublicKey

	s := 0
	for i := big.NewInt(0); cnt.Cmp(i) == 1; i.Add(i, big.NewInt(1)) {
		voter, _ := instance.Voters(nil, i)
		PP := ecdsa.PublicKey{Curve: curve, X: voter.P.X, Y: voter.P.Y}
		II := ecdsa.PublicKey{Curve: curve, X: voter.I.X, Y: voter.I.Y}
		voters = append(voters, sign.Voter{P: &PP, I: &II})
		P = append(P, PP)
		I = append(I, II)

		if sign.EqualPub(priv.PublicKey, PP) {
			s = int(i.Int64())
		}
	}

	return voters, P, I, s

}

func retriveBallots(instance *contract.Contract) ([]sign.Ballot, *ecdsa.PrivateKey) {
	cnt, _ := instance.BallotCount(nil)
	curve := math.EllipticCurve

	var ballots []sign.Ballot

	for i := big.NewInt(0); cnt.Cmp(i) == 1; i.Add(i, big.NewInt(1)) {
		ballot, _ := instance.Ballots(nil, i)
		SA := ecdsa.PublicKey{Curve: curve, X: ballot.SA.X, Y: ballot.SA.Y}
		R := ecdsa.PublicKey{Curve: curve, X: ballot.R.X, Y: ballot.R.Y}
		ballots = append(ballots, sign.Ballot{SA: SA, R: R, Signature: ballot.Signature})
	}

	m1, _ := instance.Managers(nil, big.NewInt(0))
	m2, _ := instance.Managers(nil, big.NewInt(1))

	m := big.NewInt(1)
	m = m.Mul(m1.PrivateKey, m2.PrivateKey)
	mX, mY := curve.ScalarBaseMult(m.Bytes())

	pubkey := ecdsa.PublicKey{Curve: curve, X: mX, Y: mY}
	A := ecdsa.PrivateKey{PublicKey: pubkey, D: m}

	return ballots, &A
}

func tallyElection(instance *contract.Contract, contractAddress string, candidates []sign.Candidate, P []ecdsa.PublicKey) {
	phase, _ := instance.Phase(nil)
	for phase != 2 {
		phase, _ = instance.Phase(nil)
	}

	instance = contract.Connect(contractAddress)
	ballots, A := retriveBallots(instance)

	// Tally results
	votes := make([]int, len(candidates))
	for i := 0; i < len(ballots); i++ {
		ballot := ballots[i]

		for j := 0; j < len(candidates); j++ {
			message, signature := sign.BuildAssets(ballot)

			validitySA := sign.VerifySA(A, *candidates[j].Pubkey, ballot.SA, ballot.R)
			validitySignature := sign.VER([]byte(message), P, signature)

			if validitySA && validitySignature {
				votes[j] = votes[j] + 1
			}
		}
	}

	fmt.Println()
	fmt.Println("Final results")
	for i := 0; i < len(candidates); i++ {
		fmt.Printf(" (%v) %v: %v votes\n", i, candidates[i].Name, votes[i])
	}

}

func main() {
	// Connecting to blockchain
	auth, instance, address, contractAddress := contract.ConnectionCLI()

	// Generating pubkey and registering voter
	P, I := sign.GEN(math.EllipticCurve)
	_, err := instance.AddVoter(auth, P.X, P.Y, I.X, I.Y)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Get voter count
	cnt, err := instance.VoterCount(nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Voter count:", cnt)
	fmt.Println()

	candidate, candidates := requestVote(instance, address, contractAddress)

	// Get public key
	auth, _ = contract.Auth(address)
	instance = contract.Connect(contractAddress)
	manager, _ := instance.Managers(nil, big.NewInt(0))
	Ax, Ay := manager.AnnouncedKey.X, manager.AnnouncedKey.Y

	// Verify if key is valid
	_, err = instance.GetKey(auth)
	if err != nil {
		panic(err)
	}

	// Algorithm to vote
	_, Pvec, _, s := retriveParams(instance, P)

	A := &ecdsa.PublicKey{Curve: math.EllipticCurve, X: Ax, Y: Ay}
	B := candidate.Pubkey
	SA, R := sign.GenerateSA(A, B)

	message := fmt.Sprintf("(%v,%v,%v,%v)", SA.X, SA.Y, R.X, R.Y)
	signature := sign.SIG([]byte(message), P, I, Pvec, s)
	signatureString := signature.ToStringSignature()
	byteStr, _ := json.Marshal(signatureString)
	str := string(byteStr)

	// Transmit to blockchain
	auth, _ = contract.Auth(address)
	instance = contract.Connect(contractAddress)
	_, err = instance.Vote(auth, SA.X, SA.Y, R.X, R.Y, str)
	if err != nil {
		panic(err)
	}

	// Audit results
	tallyElection(instance, contractAddress, candidates, Pvec)
}
