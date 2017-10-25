package rsa

import (
  "fmt"
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
  p := big.NewInt(521) // RandomBigPrime()
  q := big.NewInt(383) // RandomBigPrime()

  n   := big.NewInt(0)

  one := big.NewInt(1)

  p_1 := big.NewInt(0)
  q_1 := big.NewInt(0)

  pq := big.NewInt(0)
  pq.Mul(p,q)

  pq_1 := big.NewInt(0)

  p_1.Sub(p,one)
  q_1.Sub(q,one)

  pq_1.Mul(p_1,q_1)

  n.Mul(p,q)

  e := big.NewInt(227) // big.NewInt((int64)(RandomPrime()))

  mdc, x, y := extendEuclid(big.NewInt(99),big.NewInt(78))

  fmt.Println(mdc,x,y)

  d := big.NewInt(0)
  d.Mod(x,pq)

  return &PrivateCert{N: n, D: d}, &PublicCert{N: n, E: e}
}
