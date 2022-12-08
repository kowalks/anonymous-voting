package main

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/kowalks/anonymous-voting/pkg/math"
	"github.com/kowalks/anonymous-voting/pkg/sign"
)

func main() {
	N := 5
	curve := math.EllipticCurve

	key := make([]ecdsa.PrivateKey, N)
	pub := make([]ecdsa.PublicKey, N)
	img := make([]ecdsa.PublicKey, N)
	for i := 0; i < N; i++ {
		key[i], img[i] = sign.GEN(curve)
		pub[i] = key[i].PublicKey
	}

	m := []byte("Message")
	for s := 0; s < N; s++ {

		I, c, r := sign.SIG(m, key[s], img[s], pub, s)

		ver := sign.VER(m, pub, I, c, r)
		fmt.Println(ver)
	}

}
