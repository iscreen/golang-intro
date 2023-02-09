package main

import (
	"fmt"
	"sort"
)

// START NewSort OMIT
func NewSort[T string](data []T) { // HL12
	sort.Slice(data, func(i, j int) bool {
		return data[i] > data[j]
	})
}

// END NewSort OMIT

func main() {
	fruits := []string{"peach", "banana", "kiwi"}
	NewSort(fruits)
	fmt.Print(fruits)
}
