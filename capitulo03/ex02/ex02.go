package main

import (
	"fmt"
)

var x int
var y string
var z bool

func main() {
	fmt.Printf("%v %T\n", x, x)
	fmt.Printf("%v %T\n", y, y)
	fmt.Printf("%v %T\n", z, z)
}
