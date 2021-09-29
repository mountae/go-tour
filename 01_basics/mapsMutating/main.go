package main

import "fmt"

func main() {
	m := make(map[string]int)

	// insert or update an element in map 'm'.
	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	// Test that a key is present with a two-value assignment: elem, ok = m[key]
	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)

	// If 'elem' or 'ok' have not yet been declared you could use a short declaration form:
	// elem, ok := m[key]
}
