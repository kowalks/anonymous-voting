package sign

import "crypto/ecdsa"

func EqualPub(p1, p2 ecdsa.PublicKey) bool {
	if p1.Curve != p2.Curve {
		return false
	}

	cmpX := p1.X.Cmp(p2.X)
	cmpY := p1.Y.Cmp(p2.Y)

	return cmpX == 0 && cmpY == 0
}

func EqualKey(p1, p2 ecdsa.PrivateKey) bool {
	if p1.D != p2.D {
		return false
	}

	return EqualPub(p1.PublicKey, p2.PublicKey)
}
