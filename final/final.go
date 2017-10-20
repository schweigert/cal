package main

import "fmt"

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

func main() {

}
