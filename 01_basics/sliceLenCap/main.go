package main

import "fmt"

// The length of a slice is the number of elements it contains.
// The capacity of a slice is the number of elements in the underlying array,
// counting from the first element in the slice.

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)

	// Nil slice
	var ns []int
	fmt.Println(ns, len(ns), cap(ns))
	if ns == nil {
		fmt.Println("nil!")
	}

	// s := []int{2, 3, 5, 7, 11, 13}

	// s = s[1:4]
	// fmt.Println(s)

	// s = s[:2]
	// fmt.Println(s)

	// s = s[1:]
	// fmt.Println(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
