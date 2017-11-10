package rsa

import (
  "math/big"
  "fmt"
)

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

  fmt.Println(p)
  fmt.Println(q)

  n := Mul(p,q)

  phi := Mul(Sub(p,NewNum(1)),Sub(q,NewNum(1)))

  e := randomSmallPrime()

  d := multiplicativeInverse(e,phi)

  return &PrivateCert{N: n, D: d}, &PublicCert{N: n, E: e}
}
