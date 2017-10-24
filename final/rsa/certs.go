package rsa

import (
  "math/big"
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
  p := RandomBigPrime()
  q := RandomBigPrime()

  one := big.NewInt(1)

  n := big.NewInt(0)
  e := big.NewInt(0)
  p_1 := big.NewInt(0)
  q_1 := big.NewInt(0)

  p_1.Sub(p, one)
  q_1.Sub(q, one)

  n.Mul(p,q)

  return &PrivateCert{N: n}, &PublicCert{N: n}
}
