package sign

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/json"
	"fmt"

	"github.com/kowalks/anonymous-voting/pkg/math"
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

	return EqualPub(X, SA)
}

func BuildAssets(ballot Ballot) (string, Signature) {
	message := fmt.Sprintf("(%v,%v,%v,%v)", ballot.SA.X, ballot.SA.Y, ballot.R.X, ballot.R.Y)
	stringSignature := StringSignature{}
	json.Unmarshal([]byte(ballot.Signature), &stringSignature)
	signature := stringSignature.ToSignature()

	return message, signature
}
