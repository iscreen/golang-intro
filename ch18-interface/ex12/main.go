package main

import "fmt"

// START animal OMIT
type animal interface {
	breathe() // HL12
	walk()    // HL12
}

// END animal OMIT

// START mammal OMIT
type mammal interface {
	feed() // HL12
}

// END mammal OMIT

// START lion OMIT
type lion struct {
	age int
}

// END lion OMIT

// START lion methods OMIT
func (l lion) breathe() {
	fmt.Println("Lion breathes")
}

func (l lion) walk() {
	fmt.Println("Lion walk")
}

func (l lion) feed() {
	fmt.Println("Lion feeds young")
}

// END lion methods OMIT

func main() {
	var a animal // HL12
	l := lion{}
	a = l
	a.breathe()
	a.walk()

	var m mammal // HL12
	m = l
	m.feed()
}
