package main

import (
	"fmt"
)

const (
	a = iota + 2018
	b
	c
	d
)

func main() {
	fmt.Printf("%d\t %d\t %d\t %d\t", a, b, c, d)
}
