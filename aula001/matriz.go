package main

import "fmt"

func newMatrix(n int) [][]int {
  r := make([][]int, n)

  for i := 0; i < n; i++ {
    r[i] = make([]int, n)
  }

  return r
}

func mullMatrix(a [][]int, b [][]int) [][]int {
  r := newMatrix(3)
  return r
}

func printMatrix(matrix [][]int) {
  for i := range matrix {
    for j := range matrix {
      fmt.Printf("%d ", matrix[i][j])
    }
    fmt.Printf("\n")
  }
  fmt.Printf("\n")
}

func main() {
  var a = newMatrix(3)
  printMatrix(a)
}
