package main

import (
  "./rsa"
  "fmt"
  "math/big"
)

func main() {

  rsa.LEN_CRIVO = 10000
  rsa.LEN_BIG_PRIME = 255

  // O(n² + Klog(K)²)
  priv, pub := rsa.NewCert()
  fmt.Println(priv)
  fmt.Println(pub)

  message := "hehe"

  encrypted_msg := priv.Encrypt(message)
  fmt.Println(encrypted_msg)

  decrypted_msg := pub.Decrypt(encrypted_msg)
  fmt.Println(decrypted_msg)

  // o(n²log(n))
  fmt.Println(bruteForce(pub))
}

// O(n²k)
// n := p*q (da chave). BEM GRANDE
// k := número de bits de n
// k := log(p*q) = log(n)
// o(n²log(n))
func bruteForce(cert *rsa.PublicCert)(*big.Int,*big.Int){
  n := cert.N
  one := rsa.NewNum(2)
  for i := rsa.NewNum(3); rsa.CompLess(i,n); i = rsa.Add(i, one) {
    for j := i; ; j = rsa.Add(j, one) {
      mul := rsa.Mul(i,j)
      if rsa.CompBigger(mul,n){
        break
      }
      if rsa.Comp(n, mul) {
        return i, j;
      }
    }
  }

  return n, rsa.NewNum(1);
}
