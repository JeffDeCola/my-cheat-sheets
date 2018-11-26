package main

import (
	"fmt"
)

type ifacer interface {
	getSomeField() string
	setSomeField(string)
}

type myStruct struct {
	someField string
}

func (i myStruct) getSomeField() string {
	return i.someField
}

func (i *myStruct) setSomeField(newValue string) {
	i.someField = newValue
}

func main() {

	// DECLARE & ASSIGN (INITIALIZE)
	i := myStruct{"Hello"}
	// Initialize the interface -  Declare and assign the proper struct
	var a ifacer = &i

	// Using the interface
	a.setSomeField("World")
	fmt.Println(a.getSomeField())
}
