package math

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"math/big"
)

// a cryptographic hash function which maps a binary sequence with arbitrary length to a finite field F module q
func HashS(hash []byte) *big.Int {
	x := big.NewInt(0)
	x = x.SetBytes(hash)
	x = x.Mod(x, EllipticCurve.Params().N)
	return x
}

// a cryptographic hash function which maps finite field points on an elliptic curve to themselves
func HashP(point ecdsa.PublicKey) ecdsa.PublicKey {
	curve := point.Curve
	params := curve.Params()
	x, y := curve.Add(point.X, point.Y, params.Gx, params.Gy)
	return ecdsa.PublicKey{Curve: curve, X: x, Y: y}
}

// a cryptographic hash function which maps a binary sequence with arbitrary length to a point on an elliptic curve
func HashB(hash []byte) ecdsa.PublicKey {
	curve := EllipticCurve
	x, y := elliptic.Unmarshal(curve, hash)
	return ecdsa.PublicKey{Curve: curve, X: x, Y: y}
}

func HashT(point ecdsa.PublicKey) []byte {
	curve := EllipticCurve
	return elliptic.Marshal(curve, point.X, point.Y)
}
