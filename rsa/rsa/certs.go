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

  p := randomBigPrime()
  q := randomBigPrime()

  n := Mul(p,q)

  phi := Mul(Sub(p,NewNum(1)),Sub(q,NewNum(1)))

  e := randomSmallPrime()

  d := multiplicativeInverse(e,phi)

  return &PrivateCert{N: n, D: d}, &PublicCert{N: n, E: e}
}

func (priv *PrivateCert)Encrypt(plaintext string) []*big.Int {

  n := priv.N
  key := priv.D
  t := len(plaintext)
  ret := make([]*big.Int,t)

  for i := 0; i < t; i++ {
    c := NewNum((int64)(plaintext[i]))
    ret[i] = Exp(c,key,n)
  }

  return ret
}

func (publ *PublicCert)Decrypt(cipher []*big.Int) string {

  n := publ.N
  key := publ.E
  t := len(cipher)
  data := make([]byte, t)

  for i := 0; i < t; i++ {
    c := cipher[i]
    data[i] = (byte)(Exp(c,key,n).Int64())
  }

  return string(data)
}
