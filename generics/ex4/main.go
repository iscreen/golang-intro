package main

import (
	"fmt"
	"sort"
)

// START NewSort OMIT
func NewSort[T string | int32](data []T) { // HL12
	sort.Slice(data, func(i, j int) bool {
		return data[i] > data[j]
	})
}

// END NewSort OMIT

func main() {
	fruits := []string{"peach", "banana", "kiwi"}
	numbers := []int32{6, 29, 39}

	NewSort(fruits)
	fmt.Println(fruits)

	NewSort(numbers)
	fmt.Println(numbers)
}
