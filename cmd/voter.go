package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"strconv"

	"github.com/kowalks/anonymous-voting/pkg/contract"
	"github.com/kowalks/anonymous-voting/pkg/math"
	"github.com/kowalks/anonymous-voting/pkg/sign"
)

func GenerateSA(A, B *ecdsa.PublicKey) (ecdsa.PublicKey, ecdsa.PublicKey) {
	curve := math.EllipticCurve
	order := curve.Params().N

	r, _ := rand.Int(rand.Reader, order)
	xR, yR := curve.ScalarBaseMult(r.Bytes())
	R := ecdsa.PublicKey{Curve: curve, X: xR, Y: yR}

	xA, yA := curve.ScalarMult(A.X, A.Y, r.Bytes())

	// SA = Hs(rA)G + B
	H := elliptic.Marshal(curve, xA, yA)
	xH, yH := curve.ScalarBaseMult(math.HashS(H).Bytes())
	xP, yP := curve.Add(xH, yH, B.X, B.Y)
	SA := ecdsa.PublicKey{Curve: curve, X: xP, Y: yP}

	return SA, R
}

func VerifySA(A *ecdsa.PrivateKey, B, SA, R ecdsa.PublicKey) bool {
	curve := math.EllipticCurve

	// x = (Hs(aR) + b) * G
	xA, yA := curve.ScalarMult(R.X, R.Y, A.D.Bytes())
	H := math.HashS(elliptic.Marshal(curve, xA, yA))

	HGx, HGy := curve.ScalarBaseMult(H.Bytes())
	xX, yX := curve.Add(HGx, HGy, B.X, B.Y)
	X := ecdsa.PublicKey{Curve: curve, X: xX, Y: yX}

	return sign.EqualPub(X, SA)
}

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

		candidates = append(candidates, sign.Candidate{name, &ecdsa.PublicKey{math.EllipticCurve, x, y}})
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
		PP := ecdsa.PublicKey{curve, voter.P.X, voter.P.Y}
		II := ecdsa.PublicKey{curve, voter.I.X, voter.I.Y}
		voters = append(voters, sign.Voter{&PP, &II})
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
		SA := ecdsa.PublicKey{curve, ballot.SA.X, ballot.SA.Y}
		R := ecdsa.PublicKey{curve, ballot.R.X, ballot.R.Y}
		ballots = append(ballots, sign.Ballot{SA, R, ballot.Signature})
	}

	manager, _ := instance.Managers(nil, big.NewInt(0))
	pubkey := ecdsa.PublicKey{curve, manager.Pubkey.X, manager.Pubkey.Y}
	A := ecdsa.PrivateKey{pubkey, manager.PrivateKey}

	return ballots, &A
}

func tallyElection(instance *contract.Contract, contractAddress string, candidates []sign.Candidate) {
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
			if VerifySA(A, *candidates[i].Pubkey, ballot.SA, ballot.R) {
				votes[j] = votes[j] + 1
			}
		}
	}

	fmt.Println()
	fmt.Println("Final results")
	for i := 0; i < len(candidates); i++ {
		fmt.Println(" (%v) %v: %v votes", i, candidates[i].Name, votes[i])
	}

}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Provide a name for this key voter")
		os.Exit(0)
	}

	// Generating random number for voter
	name := os.Args[1]

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

	A := &ecdsa.PublicKey{math.EllipticCurve, Ax, Ay}
	B := candidate.Pubkey
	SA, R := GenerateSA(A, B)

	message := fmt.Sprintf("(%v,%v)", SA, R)
	signature := sign.SIG([]byte(message), P, I, Pvec, s)
	signatureString := fmt.Sprintf("%v", signature)

	// Transmit to blockchain
	auth, _ = contract.Auth(address)
	instance = contract.Connect(contractAddress)
	_, err = instance.Vote(auth, SA.X, SA.Y, R.X, R.Y, signatureString)
	if err != nil {
		panic(err)
	}

	// Audit results
	tallyElection(instance, contractAddress, candidates)

	// ver := VerifySA(A, B, SA, R)
	// fmt.Println(ver)

	// fmt.Println(math.EllipticCurve.Params())
	// fmt.Println(math.EllipticCurve.Params().P)

	_ = name
}
