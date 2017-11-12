package rsa

import(
  "time"
  "math/big"
  "math/rand"
)


var LEN_CRIVO int
var LEN_BIG_PRIME int

// O(Klog(K))
func multiplicativeInverse(e,phi *big.Int) *big.Int {
  d := NewNum(0)
  x1 := NewNum(0)
  x2 := NewNum(1)
  y1 := NewNum(1)

  temp_phi := phi
  // O(Klog(K)²)
  for CompBiggerZero(e) {
    temp1 := Div(temp_phi,e)
    temp2 := Sub(temp_phi,Mul(temp1, e))
    temp_phi = e
    e = temp2
    // O(K)
    x := Sub(x2,Mul(temp1, x1))
    y := Sub(d,Mul(temp1, y1))

    x2 = x1
    x1 = x
    d = y1
    y1 = y

  }

  return Add(d,phi)
}

// O(n²)
func randomPrime() int {
  // O(n²)
  c := crivo(LEN_CRIVO)
  s := len(c)

  // O(1)
  s1 := rand.NewSource(time.Now().UnixNano())
  r1 := rand.New(s1)

  // O(1)
  r := r1.Intn(s/2) + s/2

  return c[r]
}

// O(n²)
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

// O(n²)
func crivo(n int) []int {
  table := make([]bool, n)

  // Inicia a table
  // O(n)
  for i := 0; i < len(table); i++ {
    table[i] = true
  }

  // 0 e 1 não são primos
  table[0] = false
  table[1] = false

  // Quantos primos encontrou
  np := 0

  // O(n²)
  for i := 0; i < len(table); i++ {
    if !table[i]{
      continue
    }

    np++
    // O(n)
    for j := i+1; j < len(table); j++ {
      if j%i == 0 {
        table[j] = false
      }
    }
  }

  ret := make([]int, np)
  np = 0
  // O(n)
  for i := 0; i < len(table); i++ {
    if table[i] {
      ret[np] = i
      np++
    }
  }

  return ret
}

// O(k)
func randomBigPrime() *big.Int {
  bits := LEN_BIG_PRIME
  for {
    // O(1)
    zero := big.NewInt(0)
    one  := big.NewInt(1)
    two  := big.NewInt(2)

    // O(k)
    r   := randomBigNumber(bits)
    r_1 := big.NewInt(0)
    // O(k)
    r_1.Sub(r, one)
    // O(klog(k)log(log(k)))
    zero.Exp(two, r_1, r)
    // O(k)
    if one.Cmp(zero) == 0 {
      return r
    }
  }
}

// O(k)
func randomBigNumber(bits int) *big.Int {
  var str string
  // O(1)
  s1 := rand.NewSource(time.Now().UnixNano())
  r1 := rand.New(s1)
  // O(k)
  for i := 0; i < bits; i++ {
    if r1.Intn(2) == 0 {
      str = str + "1"
    } else {
      str = str + "0"
    }
  }

  // O(k)
  ret := new(big.Int)
  ret.SetString(str, 2)
  return ret
}
