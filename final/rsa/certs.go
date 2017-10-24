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

  n   := big.NewInt(0)

  one := big.NewInt(1)

  p_1 := big.NewInt(0)
  q_1 := big.NewInt(0)

  pq_1 := big.NewInt(0)

  p_1.Sub(p,one)
  q_1.Sub(q,one)

  pq_1.Mul(p_1,q_1)

  n.Mul(p,q)

  e := big.NewInt((int64)(RandomPrime()))

  return &PrivateCert{N: n}, &PublicCert{N: n, E: e}
}
