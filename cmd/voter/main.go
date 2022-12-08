package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"

	"github.com/kowalks/anonymous-voting/pkg/math"
	"github.com/kowalks/anonymous-voting/pkg/sign"
)

func GenerateSA(A, B *ecdsa.PrivateKey) (ecdsa.PublicKey, ecdsa.PublicKey) {
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

func VerifySA(A, B *ecdsa.PrivateKey, SA, R ecdsa.PublicKey) bool {
	curve := math.EllipticCurve

	// x = Hs(aR) + b
	xA, yA := curve.ScalarMult(R.X, R.Y, A.D.Bytes())
	H := math.HashS(elliptic.Marshal(curve, xA, yA))
	x := H.Add(H, B.D)

	// xG = SA
	xX, yX := curve.ScalarBaseMult(x.Bytes())
	X := ecdsa.PublicKey{Curve: curve, X: xX, Y: yX}

	return sign.EqualPub(X, SA)
}

func main() {
	A, _ := ecdsa.GenerateKey(math.EllipticCurve, rand.Reader)
	B, _ := ecdsa.GenerateKey(math.EllipticCurve, rand.Reader)
	SA, R := GenerateSA(A, B)

	ver := VerifySA(A, B, SA, R)
	fmt.Println(ver)

	fmt.Println(math.EllipticCurve.Params())
	fmt.Println(math.EllipticCurve.Params().P)
}
