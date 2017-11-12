package main

import (
  "./rsa"
  "fmt"
  "math/big"
)

func main() {

  rsa.LEN_CRIVO = 10000
  rsa.LEN_BIG_PRIME = 32

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
// o(n²log(n))
func bruteForce(cert *rsa.PublicCert)(*big.Int,*big.Int){
  n := cert.N
  i := rsa.NewNum(2)
  one := rsa.NewNum(1)
  s := rsa.Sub(n,one)

  lista := make([]*big.Int, 5000)
  l := 0
  // O(nk)
  for {
    // O(k)
    if(rsa.CompBigger(i,s)){
      break
    }
    // O(k)
    if(rsa.CompEqualZero(rsa.Mod(n,i))) {
      lista[l] = i
      l++
    }
    // O(k)
    i = rsa.Add(i, one)
  }

  var a,b *big.Int
  // O(n²k)
  for m := 0; m < l; m++ {
    // O(nk)
    for k := m; k < l; k++ {
      // O(k)
      if rsa.Comp(n,rsa.Mul(lista[m],lista[k])) {
        a,b =  lista[m],lista[k]
      }
    }
  }

  return a,b
}
