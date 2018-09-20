package main

import "fmt"

func sum(n ...int) int {
	total := 0
	for _, f := range n {
		total += f
	}
	return total
}

func main() {
	data := []int{43, 44, 55, 66}
	fmt.Println(sum(data...))
}
