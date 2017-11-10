package rsa

import(
  "time"
  "math/big"
  "math/rand"
)

const(
  LEN_CRIVO = 10000
  LEN_BIG_PRIME = 32
)

func multiplicativeInverse(e,phi *big.Int) *big.Int {
  d := NewNum(0)
  x1 := NewNum(0)
  x2 := NewNum(1)
  y1 := NewNum(1)

  temp_phi := phi

  for CompBiggerZero(e) {
    temp1 := Div(temp_phi,e)
    temp2 := Sub(temp_phi,Mul(temp1, e))
    temp_phi = e
    e = temp2

    x := Sub(x2,Mul(temp1, x1))
    y := Sub(d,Mul(temp1, y1))

    x2 = x1
    x1 = x
    d = y1
    y1 = y

  }
  
  return Add(d,phi)
}

func randomPrime() int {
  c := crivo(LEN_CRIVO)
  s := len(c)


  s1 := rand.NewSource(time.Now().UnixNano())
  r1 := rand.New(s1)

  r := r1.Intn(s/2) + s/2

  return c[r]
}

func randomSmallPrime() *big.Int {
  return NewNum((int64)(randomPrime()))
}

func gcd(a, b *big.Int) *big.Int {
  for CompNotEqualZero(b) {
    c := a
    d := b
    b = Mod(c,d)
    a = d
  }
  return a
}

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

func randomBigPrime() *big.Int {
  bits := LEN_BIG_PRIME
  for {
    zero := big.NewInt(0)
    one  := big.NewInt(1)
    two  := big.NewInt(2)

    r   := randomBigNumber(bits)
    r_1 := big.NewInt(0)
    r_1.Sub(r, one)
    zero.Exp(two, r_1, r)
    if one.Cmp(zero) == 0 {
      return r
    }
  }
}

func randomBigNumber(bits int) *big.Int {
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
