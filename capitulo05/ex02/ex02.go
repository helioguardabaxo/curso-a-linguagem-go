package main

import (
	"fmt"
)

func main() {
	a := (10 == 10)
	b := (10 != 10)
	c := (5 <= 10)
	d := (5 >= 10)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
