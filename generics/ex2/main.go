package main

import "math/bits"

// START Interface OMIT

type Interface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

// END Interface OMIT

// START Sort OMIT
func Sort(data Interface) {
	n := data.Len()
	if n <= 1 {
		return
	}
	limit := bits.Len(uint(n))
	pdqsort(data, 0, n, limit)
}

// END Sort OMIT
