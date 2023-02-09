package main

import "fmt"

// START Describer OMIT
type Describer interface {
	Describe()
}

// END Describer OMIT

// START Person OMIT
type Person struct {
	name string
	age  int
}

func (p Person) Describe() { //implemented using value receiver
	fmt.Printf("%s is %d years old\n", p.name, p.age)
}

// END Person OMIT

// START Address OMIT
type Address struct {
	state   string
	country string
}

func (a *Address) Describe() { //implemented using pointer receiver // HL12
	fmt.Printf("State %s Country %s", a.state, a.country) // HL12
} // HL12

// END Address OMIT

func main() {
	var d1 Describer

	p1 := Person{"Sam", 25} // Value Type
	d1 = p1
	d1.Describe()
	p2 := Person{"James", 32}
	d1 = &p2 // Pointer Type
	d1.Describe()

	var d2 Describer
	a := Address{"Washington", "USA"} // Value Type

	// Error => Describe method has pointer receiver
	//d2 = a // HL12

	// This works since Describer interface // HL12
	d2 = &a // HL12
	// is implemented by Address pointer in line 30
	d2.Describe()

}
