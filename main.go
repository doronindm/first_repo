package main

import "fmt"

func swap(x, y *int) {
	*x, *y = *y, *x
}
func main() {
	x := 3
	y := 5
	fmt.Println("Before: ", "x =", x, "y = ", y)
	swap(&x, &y)
	fmt.Println("After: ", "x =", x, "y = ", y)
}
