package main

import "fmt"

// START Donald OMIT
type Donald struct{}

// END Donald OMIT

// START Donald Method OMIT
// Make Quack a method of Donald
func (d Donald) Quack() { //HL12
	fmt.Println("I am Donald Duck!")
}

// Make Walk a method of Donald
func (d Donald) Walk() { //HL12
	fmt.Println("I waddle")
}

// END Donald Method OMIT
func main() {
	d := new(Donald)
	d.Quack()
	d.Walk()
}
