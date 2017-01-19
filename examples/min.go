package main

import (
	"fmt"
	"math"
)

func min(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

func min2(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println("vim-go")
	if true {

	}
}
