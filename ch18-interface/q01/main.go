package main

import "fmt"

type myString string

func (s myString) String() string {
	return fmt.Sprintf("This is %s", s)
}

func main() {
	var value myString = "test"
	var s fmt.Stringer
	s = value
	fmt.Println(s)
}
