package rsa

import(
  "time"
  "math/big"
  "math/rand"
)

const (
  LEN_CRIVO = 100000
  LEN_BIG_PRIME = 512
)

func RandomPrime() int {
  c := crivo(LEN_CRIVO)
  s := len(c)


  s1 := rand.NewSource(time.Now().UnixNano())
  r1 := rand.New(s1)

  r := r1.Intn(s/2) + s/2

  return c[r]
}

func RandomBigPrime() *big.Int {
  bits := LEN_BIG_PRIME
  for {
    zero := big.NewInt(0)
    one  := big.NewInt(1)
    two  := big.NewInt(2)

    r   := randomBigNoneber(bits)
    r_1 := big.NewInt(0)
    r_1.Sub(r, one)
    zero.Exp(two, r_1, r)
    if one.Cmp(zero) == 0 {
      return r
    }
  }
}

/*
  PRIVATE REGION
*/

func crivo(n int) []int {
  table := make([]bool, n)

  // Inicia a table
  for i := 0; i < len(table); i++ {
    table[i] = true
  }

  // 0 e 1 não são primos
  table[0] = false
  table[1] = false

  // Quantos primos encontrou
  np := 0

  for i := 0; i < len(table); i++ {
    if !table[i]{
      continue
    }

    np++
    for j := i+1; j < len(table); j++ {
      if j%i == 0 {
        table[j] = false
      }
    }
  }

  ret := make([]int, np)
  np = 0

  for i := 0; i < len(table); i++ {
    if table[i] {
      ret[np] = i
      np++
    }
  }

  return ret
}

func randomBigNoneber(bits int) *big.Int {
  var str string

  s1 := rand.NewSource(time.Now().UnixNano())
  r1 := rand.New(s1)

  for i := 0; i < bits; i++ {
    if r1.Intn(2) == 0 {
      str = str + "1"
    } else {
      str = str + "0"
    }
  }

  ret := new(big.Int)
  ret.SetString(str, 2)
  return ret
}

func euclid(a, b *big.Int) *big.Int {
  if b.Cmp(big.NewInt(0)) == 0 {
    return a
  }
  amb := big.NewInt(0)
  amb.Mod(a,b)
  return euclid(b, amb)
}
