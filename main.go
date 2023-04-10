package main

import "fmt"

func up(x ...int) int {
	Max := x[0]
	for _, y := range x {
		if y > Max {
			Max = y
		}
	}
	return Max
}

func main() {
	x := []int{
		98, 73, 86, 92,
		79, 54, 21, 83,
		17, 12, 129, 36,
	}

	fmt.Println(up(x...))
}
