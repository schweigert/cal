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

  return &PrivateCert{}, &PublicCert{}
}
