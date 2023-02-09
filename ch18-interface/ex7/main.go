package main

import "fmt"

type People struct {
	name string
}

func main() {
	man := People{ // HL12
		name: "Dean", // HL12
	} // HL12

	fmt.Println("This is value type")
	fmt.Printf("man type: %T, value: %v\n", man, man)
}
