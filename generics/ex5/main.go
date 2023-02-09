package main

import (
	"fmt"
	"sort"
)

// START SortData OMIT
type SortData interface { // HL12
	string | int32
}

// END SortData OMIT
// START NewSort OMIT
func NewSort[T SortData](data []T) { // HL12
	sort.Slice(data, func(i, j int) bool {
		return data[i] > data[j]
	})
}

// END NewSort OMIT

func main() {
	fruits := []string{"peach", "banana", "kiwi"}
	numbers := []int32{6, 29, 39}

	NewSort(fruits)
	fmt.Print(fruits)
	NewSort(numbers)
	fmt.Println(numbers)
}
