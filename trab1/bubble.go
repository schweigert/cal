package main

import "fmt"

func main() {
	var size int;
	fmt.Scanf("%d", &size)
	var source []int = make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Scanf("%d", &source[i])
	}

	// Bubble

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if source[i] < source[j] {
				 source[i], source[j] = source[j], source[i]
			}
		}
	}
}
