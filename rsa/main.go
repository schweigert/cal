package main

import (
  "./rsa"
  "fmt"
  "math/big"
)

func main() {

  rsa.LEN_CRIVO = 10000
  rsa.LEN_BIG_PRIME = 32

  priv, pub := rsa.NewCert()
  fmt.Println(priv)
  fmt.Println(pub)

  message := "hehe"

  encrypted_msg := priv.Encrypt(message)
  fmt.Println(encrypted_msg)

  decrypted_msg := pub.Decrypt(encrypted_msg)
  fmt.Println(decrypted_msg)

  fmt.Println(bruteForce(pub))
}


func bruteForce(cert *rsa.PublicCert)(*big.Int,*big.Int){
  n := cert.N
  i := rsa.NewNum(2)
  one := rsa.NewNum(1)
  s := rsa.Sub(n,one)

  lista := make([]*big.Int, 300)
  l := 0

  for {
    if(rsa.CompBigger(i,s)){
      break
    }

    if(rsa.CompEqualZero(rsa.Mod(n,i))) {
      lista[l] = i
      l++
    }

    i = rsa.Add(i, one)
  }

  var a,b *big.Int

  for m := 0; m < l; m++ {
    for k := m; k < l; k++ {
      if rsa.Comp(n,rsa.Mul(lista[m],lista[k])) {
        a,b =  lista[m],lista[k]
      }
    }
  }

  return a,b
}
