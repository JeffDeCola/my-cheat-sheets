package main

import "fmt"

// Returns a function
func inc(x int) func() int {
	return func() int {
		x++
		return x
	}
}

func main() {
	a := [2]float32{1.1, 2.0} // Array Shorthand Assignment
	fmt.Println(a, len(a), cap(a))
}
