package main

import "fmt"

// START Empty Interface OMIT
type any = interface{} // HL12

func displayValue(i any) {
	fmt.Printf("type: %T, value: %v\n", i, i)
}

// END Empty Interface OMIT

func main() {

	a := "Welcome to Programiz"
	b := 20
	c := false

	// pass string to the function
	displayValue(a)

	// pass integer number to the function
	displayValue(b)

	// pass boolean value to the function
	displayValue(c)

}
