package math

import (
	"crypto/elliptic"
	"fmt"
)

var EllipticCurve = elliptic.P224()

func PrintCurve(c elliptic.Curve) {
	params := c.Params()

	fmt.Println("Curve parameters")
	fmt.Println("    Name:   ", params.Name)
	fmt.Println("    Prime:  ", params.P)
	fmt.Println("    G order:", params.N)
	fmt.Println("    G = ", params.Gx, params.Gy)
	fmt.Println()
}
