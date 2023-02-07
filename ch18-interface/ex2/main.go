package main

import "fmt"

// START IronMan OMIT
type IronMan struct{}

// END IronMan OMIT

// START IronMan Method OMIT
func (t IronMan) Quack() {
	fmt.Println("I am Iron man!")
}
func (t IronMan) Walk() {
	fmt.Println("I dislike walk, I can fly")
}

// END IronMan Method OMIT

func main() {
	i := new(IronMan)
	i.Quack()
	i.Walk()
}
