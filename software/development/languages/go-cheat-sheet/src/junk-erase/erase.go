package main

import "fmt"

// Passing a function - callback
func math(numbers []int, s int, f func(int, int)) {
	for _, n := range numbers {
		f(n, s) // Use the function passed from main
		fmt.Println("math ", n, s)
	}
}

// Returning a function

func main() {

	numbers := []int{1, 2, 3, 4}
	sum := 0

	// func as argument - Print slice
	math(numbers, sum, func(n int, s int) {
		s = n
		fmt.Println("just same", n, s)

	})

	// func as argument - Print sum
	math(numbers, sum, func(n int, s int) {
		s = s + n
		fmt.Println("sum ", n, s)

	})

	// func as argument - Print average
	math(numbers, sum, func(n int, s int) {
		s = s + n
		s = s / 2
		fmt.Println("average ", n, s)

	})
}
