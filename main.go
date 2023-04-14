package main

import "fmt"

type animal interface {
	swimming
	diving
	flight
}
type swimming interface {
	swim()
}
type diving interface {
	dive()
}
type flight interface {
	fly()
}
type flyingFish struct{}

func (f *flyingFish) swim() {
	fmt.Println("Flying fish swim")
}
func (f *flyingFish) dive() {
	fmt.Println("Flying fish dive")
}
func (f *flyingFish) fly() {
	fmt.Println("Flying fish fly")
}
func main() {
	var f animal = &flyingFish{}
	f.swim()
	f.dive()
	f.fly()
}
