package main

import "fmt"

type fish interface {
	swimming
	diving
	flight
}
type bird interface {
	walker
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
type walker interface {
	walk()
}
type flyingFish struct{}
type seagull struct{}

func (f *flyingFish) swim() {
	fmt.Println("Flying fish swim")
}
func (f *flyingFish) dive() {
	fmt.Println("Flying fish dive")
}
func (f *flyingFish) fly() {
	fmt.Println("Flying fish fly")
}
func (s *seagull) walk() {
	fmt.Println("Seagull walk")
}
func (s *seagull) dive() {
	fmt.Println("Seagull dive")
}
func (s *seagull) fly() {
	fmt.Println("Seagull fly")
}
func dive(d diving) {
	d.dive()
}
func main() {
	var f fish = &flyingFish{}
	var s bird = &seagull{}
	dive(f)
	dive(s)
}
