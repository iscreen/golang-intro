package main

import "fmt"

// START animal OMIT
type animal interface {
	breathe() // HL12
	walk()    // HL12
}

// END animal OMIT

// START human OMIT
type human interface {
	animal //embedding // HL12
	speak()
}

// END human OMIT

// START staff OMIT
type staff struct { // HL12
	name string // HL12
} // HL12

func (s staff) breathe() { // HL12
	fmt.Printf("Staff %s can breathe\.n", s.name)
}

func (s staff) walk() { // HL12
	fmt.Printf("Staff %s can walk.\n", s.name)
}

func (s staff) speak() { // HL12
	fmt.Printf("Staff %s can speak.\n", s.name)
}

// END staff OMIT

func main() {
	var h human // HL12

	h = staff{name: "Jeffery"}

	h.breathe()
	h.walk()
	h.speak()
}
