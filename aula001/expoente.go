package main

import "fmt"

func expr(a int, b int) int {
	r := 1
	for b != 0 {
		r *= a
		b--
	}
	return r
}

func main() {
	var a,b int
	fmt.Scanf("%d %d", &a, &b)
	fmt.Printf("\n%d\n", expr(a,b))
}
