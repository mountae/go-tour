package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func swap(c, d string) (string, string) {
	return d, c
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(add(42, 13))
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	fmt.Println(split(17))
}
