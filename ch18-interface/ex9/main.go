package main

import "fmt"

type People struct {
	name string
}

// START Value Receiver OMIT
func (p People) Name() string { // HL12
	p.name = "Dean Lin"
	return p.name // HL12
} // HL12

// END Value Receiver OMIT

func main() {
	man := People{
		name: "Dean",
	}

	fmt.Println("This is value receiver")
	fmt.Printf("man's name is: %s, value: %v\n", man.Name(), man)
}
