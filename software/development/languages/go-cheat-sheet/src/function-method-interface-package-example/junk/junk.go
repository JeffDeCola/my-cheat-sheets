package main

import (
	"fmt"
)

type i interface {
	Host()
}

type Host struct {
	Name string
}

func main() {

	//var i interface{ Host() }
	var a i = Host{"jeff"}

	// Map["hosts"] = []Host{Host{"test.com"}, Host{"test2.com"}}

	// hm := Map["hosts"].([]Host)
	fmt.Println(t)
}
