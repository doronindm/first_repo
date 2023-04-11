package main

import "fmt"

func up(x ...int) int {
	Min := x[0]
	for _, y := range x {
		if y < Min {
			Min = y
		}
	}
	return Min
}

func main() {
	x := []int{
		98, 73, 86, 92,
		79, 54, 21, 83,
		17, 12, 129, 36,
	}

	fmt.Println(up(x...))
}
