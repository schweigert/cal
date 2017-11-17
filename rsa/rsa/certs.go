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

// O(n² + Klog(K)²)
// n := Tamanho do Crivo
// k := número de Bits
func NewCert() (*PrivateCert, *PublicCert) {
  // O(k)
  p := randomBigPrime()
  q := randomBigPrime()
  fmt.Println(p,q)
  // O(klog(k)log(log(k)))
  n := Mul(p,q)

  // O(k)
  phi := Mul(Sub(p,NewNum(1)),Sub(q,NewNum(1)))

  // O(n²)
  e := randomSmallPrime()

  // O(Klog(K)²)
  d := multiplicativeInverse(e,phi)

  return &PrivateCert{N: n, D: d}, &PublicCert{N: n, E: e}
}

// O(sk)
// S := Tamanho do texto
// K := Tamanho em bits da chave
func (priv *PrivateCert)Encrypt(plaintext string) []*big.Int {

  n := priv.N
  key := priv.D
  t := len(plaintext)
  ret := make([]*big.Int,t)

  // O(sk)
  for i := 0; i < t; i++ {
    c := NewNum((int64)(plaintext[i]))
    // O(k)
    ret[i] = Exp(c,key,n)
  }

  return ret
}
// O(nk)
// n := número de blocos
// k := número de bits da chave
func (publ *PublicCert)Decrypt(cipher []*big.Int) string {

  n := publ.N
  key := publ.E
  t := len(cipher)
  data := make([]byte, t)
  // O(n)
  for i := 0; i < t; i++ {
    c := cipher[i]
    // O(k)
    data[i] = (byte)(Exp(c,key,n).Int64())
  }

  return string(data)
}
