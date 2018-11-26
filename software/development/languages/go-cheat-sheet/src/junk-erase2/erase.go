package main

import (
	"fmt"
)

type theInterface interface {
	doThis(n int)
}

type a struct {
	name string
}

func (i *a) doThis(n int) {
	fmt.Printf("I'm in A %v - %v\n", n, i.name)
}

type b struct {
	x int
	y int
}

func (i *b) doThis(n int) {
	fmt.Printf("I'm in B %v - %v %v\n", n, i.x, i.y)
}

// The interface accepts ANYTHING as long as that
// ANYTHING has the method
func magic(i theInterface, n int) {
	i.doThis(n)
}

func main() {
	var a1 = &a{"jeff"}
	var b1 = &b{222, 333}

	// Magic uses an interface, send it ANYTHING
	// As long as it has the underlying method.
	magic(a1, 44)
	magic(b1, 111)

	a1.doThis(44)
	b1.doThis(111)

}
