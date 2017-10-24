package main

import (
  "fmt"
  "math/big"
)

const (
  bits = 16
  number_a = "1010101010101011"
  number_b = "1110110110101111"
)

func euclid(a, b *big.Int) *big.Int {
  if b.Cmp(big.NewInt(0)) == 0 {
    return a
  }
  amb := big.NewInt(0)
  amb.Mod(a,b)
  return euclid(b, amb)
}

func main() {
  a := new(big.Int)
  a.SetString(number_a, 2) // Base binária
  b := new(big.Int)
  b.SetString(number_b, 2) // Base binária
  r := euclid(a,b)
  fmt.Println(r)
}
