package main

import "fmt"

// START Duck OMIT
type Duck interface {
	Quack() // HL12
	Walk()
}

// END Duck OMIT

type Donald struct{}

func (d Donald) Quack() { //HL12
	fmt.Println("I am Donald Duck!")
}

func (d Donald) Walk() { //HL12
	fmt.Println("I waddle")
}

type IronMan struct{}

func (t IronMan) Quack() {
	fmt.Println("I am Iron man!")
}
func (t IronMan) Walk() {
	fmt.Println("I dislike walk, I can fly")
}

// START behaveLikeDuck OMIT
// Define function argument to be of type Duck
func behaveLikeDuck(d Duck) {
	d.Quack()
	d.Walk()
}

// END behaveLikeDuck OMIT

func main() {
	ducks := []Duck{
		new(Donald),
		new(IronMan),
	}

	for _, d := range ducks {
		behaveLikeDuck(d)
	}
}
