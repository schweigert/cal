package main

import(
  "fmt"
  "time"
  "math/big"
  "math/rand"
  )

const (
  TAM_CRIVO = 100000
)

// Retorna uma lista de números primos
func crivo(n int) []int {
  tabela := make([]bool, n)

  // Inicia a tabela
  for i := 0; i < len(tabela); i++ {
    tabela[i] = true
  }

  // 0 e 1 não são primos
  tabela[0] = false
  tabela[1] = false

  // Quantos primos encontrou
  np := 0

  for i := 0; i < len(tabela); i++ {
    if !tabela[i]{
      continue
    }

    np++
    for j := i+1; j < len(tabela); j++ {
      if j%i == 0 {
        tabela[j] = false
      }
    }
  }

  ret := make([]int, np)
  np = 0

  for i := 0; i < len(tabela); i++ {
    if tabela[i] {
      ret[np] = i
      np++
    }
  }

  return ret
}

// Retorna um int (primo) aleatório
func gera_primo_pequeno() int {
  c := crivo(TAM_CRIVO)
  s := len(c)


  s1 := rand.NewSource(time.Now().UnixNano())
  r1 := rand.New(s1)

  r := r1.Intn(s/2) + s/2

  return c[r]
}

func gera_numero_aleatorio(bits int) *big.Int {
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
  ret.SetString(str, 2) // Base binária
  return ret
}

func primo_provavel(bits int) *big.Int {
  for {
    um   := big.NewInt(1)
    dois := big.NewInt(2)

    r   := gera_numero_aleatorio(bits)
    r_1 := big.NewInt(0)
    r_1.Sub(r, um)
    z := big.NewInt(0)
    z.Exp(dois, r_1, r)
    if um.Cmp(z) == 0 {
      return r
    }
  }
}

func euclid(a, b *big.Int) *big.Int {
  if b.Cmp(big.NewInt(0)) == 0 {
    return a
  }
  amb := big.NewInt(0)
  amb.Mod(a,b)
  return euclid(b, amb)
}

func euclid_ext(a, b *big.Int) (*big.Int, *big.Int, *big.Int) {
  if b.Cmp(big.NewInt(0)) == 0 {
    return a, big.NewInt(1), big.NewInt(0)
  }
  r := big.NewInt(0)
  r.Mod(a,b)
  dl, xl, yl := euclid_ext(b, r)
  y := big.NewInt(0)
  div := big.NewInt(0)
  mul := big.NewInt(0)
  div.Div(a,b)
  mul.Mul(div,yl)
  y.Sub(xl, mul)
  return dl, xl, y
}

func main() {
  fmt.Println(euclid_ext(primo_provavel(1024),primo_provavel(1024)))
}
