package rsa

import "math/big"

type PublicCert struct {
  E *big.Int
  N *big.Int
}

type PrivateCert struct {
  D *big.Int
  N *big.Int
}

func NewCert() (*PrivateCert, *PublicCert) {

  p := randomBigPrime()
  q := randomBigPrime()

  n := Mul(p,q)

  phi := Mul(Sub(p,NewNum(1)),Sub(q,NewNum(1)))

  e := randomSmallPrime()

  return &PrivateCert{N: n}, &PublicCert{N: n, E: e}

}
