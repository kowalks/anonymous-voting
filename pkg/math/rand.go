package math

import (
	"crypto/rand"
	"math/big"
)

func GenerateRand(number string) *big.Int {
	rnd := new(big.Int)
	rnd, ok := rnd.SetString(number, 10)
	if !ok {
		panic("Error in converting number to big.Int")
	}

	if len(rnd.Bytes()) == 0 {
		order := EllipticCurve.Params().N
		rnd, _ = rand.Int(rand.Reader, order)
	}

	return rnd
}
