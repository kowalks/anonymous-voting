package sign

import (
	"crypto/ecdsa"
	"math/big"
)

type Candidate struct {
	Name   string
	Pubkey *ecdsa.PublicKey
}

type Voter struct {
	P *ecdsa.PublicKey
	I *ecdsa.PublicKey
}

type Signature struct {
	I ecdsa.PublicKey
	C []*big.Int
	R []*big.Int
}

type Ballot struct {
	SA        ecdsa.PublicKey
	R         ecdsa.PublicKey
	Signature string
}
