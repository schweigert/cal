package main

import (
  "./rsa"
  "fmt"
)

func main() {
  priv, pub := rsa.NewCert()
  fmt.Println(priv)
  fmt.Println(pub)
}
