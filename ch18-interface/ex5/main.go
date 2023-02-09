package main

import (
	"fmt"
)

// START assertion OMIT
func assert(i interface{}) {
	v, ok := i.(int) // HL12
	fmt.Println(v, ok)
}

// END assertion OMIT

func main() {
	var s interface{} = 56
	assert(s)
	var i interface{} = "Steven Paul"
	assert(i)
}
