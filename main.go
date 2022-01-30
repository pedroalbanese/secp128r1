// CC0 1.0 Universal

package secp128r1

import "crypto/elliptic"
import "math/big"

func strbig(s string) (i *big.Int) {
	i = new(big.Int)
	i.SetString(s,0)
	return
}

var secp128r1 = &elliptic.CurveParams{
	P: strbig("0x00fffffffdffffffffffffffffffffffff"), // Prime
	N: strbig("0x00fffffffe0000000075a30d1b9038a115"), // Order
	B: strbig("0x00e87579c11079f43dd824993c2cee5ed3"), // B
	Gx: strbig("0x161ff7528b899b2d0c28607ca52c5b86"),  // Generator X
	Gy: strbig("0xcf5ac8395bafeb13c02da292dded7a83"),  // Generator Y
	BitSize: 128,
	Name: "secp128r1",
}

// Secp128() returns a Curve which implements Secp128r1
func Secp128r1() elliptic.Curve { return secp128r1 }



