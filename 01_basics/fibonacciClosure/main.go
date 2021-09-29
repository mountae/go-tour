package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	fib := 0
	return func() int {
		first, second := 0, 1
		for i := 0; i < fib; i++ {
			first, second = second, first+second
		}
		fib++
		return first
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
