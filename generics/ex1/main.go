package main

import (
	"fmt"
	"sort"
)

// START Interface OMIT
type ByLength []string

func (s ByLength) Len() int {
	return len(s)
}

func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

// END Interface OMIT
func main() {
	fruits := []string{"peach", "banana", "kiwi"}

	sort.Sort(ByLength(fruits)) // HL12
	fmt.Println(fruits)
}
