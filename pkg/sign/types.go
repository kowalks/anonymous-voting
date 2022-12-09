package sign

import (
	"crypto/ecdsa"
	"math/big"

	"github.com/kowalks/anonymous-voting/pkg/math"
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

type StringSignature struct {
	Ix *big.Int
	Iy *big.Int
	C  []*big.Int
	R  []*big.Int
}

type Ballot struct {
	SA        ecdsa.PublicKey
	R         ecdsa.PublicKey
	Signature string
}

func (s Signature) ToStringSignature() StringSignature {
	return StringSignature{s.I.X, s.I.Y, s.C, s.R}
}

func (s StringSignature) ToSignature() Signature {
	key := ecdsa.PublicKey{Curve: math.EllipticCurve, X: s.Ix, Y: s.Iy}
	return Signature{key, s.C, s.R}
}
