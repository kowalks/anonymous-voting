package sign

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"math/big"

	"github.com/kowalks/anonymous-voting/pkg/math"
)

func CHA(message []byte, L, R []ecdsa.PublicKey) *big.Int {
	for i := 0; i < len(L); i++ {
		message = append(message, math.HashT(L[i])...)
	}
	for i := 0; i < len(R); i++ {
		message = append(message, math.HashT(R[i])...)
	}
	return math.HashS(message)
}

func GEN(curve elliptic.Curve) (ecdsa.PrivateKey, ecdsa.PublicKey) {
	key, _ := ecdsa.GenerateKey(curve, rand.Reader)

	Hp := math.HashP(key.PublicKey)
	x, y := curve.ScalarMult(Hp.X, Hp.Y, key.D.Bytes())
	I := ecdsa.PublicKey{Curve: curve, X: x, Y: y}

	return *key, I
}

func SIG(message []byte, key ecdsa.PrivateKey, I ecdsa.PublicKey, pub []ecdsa.PublicKey, s int) Signature {
	if !EqualPub(pub[s], key.PublicKey) {
		panic("Private key not in the public list")
	}

	curve := key.Curve
	order := curve.Params().N
	N := len(pub)
	if s < 0 || s >= N {
		panic("s not in range")
	}

	L := make([]ecdsa.PublicKey, N)
	R := make([]ecdsa.PublicKey, N)
	q := make([]big.Int, N)
	w := make([]big.Int, N)

	for i := 0; i < N; i++ {
		rnd, _ := rand.Int(rand.Reader, order)
		q[i] = *rnd

		rnd, _ = rand.Int(rand.Reader, order)
		w[i] = *rnd

		if i == s {
			x, y := curve.ScalarBaseMult(q[i].Bytes())
			L[i] = ecdsa.PublicKey{Curve: curve, X: x, Y: y}

			Hp := math.HashP(pub[i])
			x, y = curve.ScalarMult(Hp.X, Hp.Y, q[i].Bytes())
			R[i] = ecdsa.PublicKey{Curve: curve, X: x, Y: y}
		} else {
			x, y := curve.ScalarBaseMult(q[i].Bytes())
			xP, yP := curve.ScalarMult(pub[i].X, pub[i].Y, w[i].Bytes())
			xL, yL := curve.Add(x, y, xP, yP)
			L[i] = ecdsa.PublicKey{Curve: curve, X: xL, Y: yL}

			Hp := math.HashP(pub[i])
			x, y = curve.ScalarMult(Hp.X, Hp.Y, q[i].Bytes())
			xI, yI := curve.ScalarMult(I.X, I.Y, w[i].Bytes())
			xR, yR := curve.Add(x, y, xI, yI)
			R[i] = ecdsa.PublicKey{Curve: curve, X: xR, Y: yR}
		}
	}

	challenge := CHA(message, L, R)

	c := make([]*big.Int, N)
	r := make([]*big.Int, N)
	for i := 0; i < N; i++ {
		if i != s {
			c[i] = &w[i]
			r[i] = &q[i]
		} else {
			c[i] = challenge.Add(challenge, &w[i])
			for j := 0; j < N; j++ {
				c[i] = c[i].Add(c[i], order)
				c[i] = c[i].Sub(c[i], &w[j])
				c[i] = c[i].Mod(c[i], order)
			}
			r[i] = big.NewInt(0)
			r[i] = r[i].Mul(c[s], key.D)
			r[i] = r[i].Sub(&q[s], r[i])
			r[i] = r[i].Add(r[i], order)
			r[i] = r[i].Mod(r[i], order)
		}
	}

	return Signature{I, c, r}
}

func VER(message []byte, pub []ecdsa.PublicKey, s Signature) bool {
	I, c, r := s.I, s.C, s.R
	curve := I.Curve
	order := curve.Params().N
	N := len(pub)

	L := make([]ecdsa.PublicKey, N)
	R := make([]ecdsa.PublicKey, N)
	sum := big.NewInt(0)
	for i := 0; i < N; i++ {
		xG, yG := curve.ScalarBaseMult(r[i].Bytes())
		xP, yP := curve.ScalarMult(pub[i].X, pub[i].Y, c[i].Bytes())
		xL, yL := curve.Add(xG, yG, xP, yP)
		L[i] = ecdsa.PublicKey{Curve: curve, X: xL, Y: yL}

		Hp := math.HashP(pub[i])
		xH, yH := curve.ScalarMult(Hp.X, Hp.Y, r[i].Bytes())
		xI, yI := curve.ScalarMult(I.X, I.Y, c[i].Bytes())
		xR, yR := curve.Add(xH, yH, xI, yI)
		R[i] = ecdsa.PublicKey{Curve: curve, X: xR, Y: yR}

		sum = sum.Add(sum, c[i])
		sum = sum.Mod(sum, order)
	}

	challenge := CHA(message, L, R)
	challenge = challenge.Mod(challenge, order)

	eq := challenge.Cmp(sum)
	return eq == 0
}
