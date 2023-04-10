package main

import "fmt"

func sum(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total
}

func main() {
	xs := []float64{
		98, 73, 86, 92,
		79, 54, 21, 83,
		17, 12, 7, 36,
	}

	fmt.Println(sum(xs))
}