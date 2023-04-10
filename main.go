package main

import (
	"fmt"
	"math"
)

func half(n int) (float64, bool) {
	halfn := float64(n) / 2
	if math.Mod(halfn, 2) == 0 {
		return halfn, true
	}
	return halfn, false
}

func main() {
	fmt.Println(half(30))
	fmt.Println(half(31))
}
