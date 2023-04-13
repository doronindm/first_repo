package main

import "fmt"

type animal interface {
	makeSound()
}
type cat struct{}
type dog struct{}

func (c *cat) makeSound() {
	fmt.Println("Meow!")
}
func (c *dog) makeSound() {
	fmt.Println("Woof!")
}
func main() {
	var c, d animal = &cat{}, &dog{}

	c.makeSound()
	d.makeSound()
}
