package main

import "fmt"

type Duck interface {
	Quack()
}

func main() {
	var duck Duck // HL12
	fmt.Printf("duck is nil and has type %T value %v\n", duck, duck)
}
