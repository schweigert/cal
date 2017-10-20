package main

import(
  "fmt"
  "time"
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

func gera_primo_pequeno() int {
  c := crivo(TAM_CRIVO)
  s := len(c)


  s1 := rand.NewSource(time.Now().UnixNano())
  r1 := rand.New(s1)

  r := r1.Intn(s/2) + s/2

  return c[r]
}

func main() {
  fmt.Println(gera_primo_pequeno())
}
