package main

import (
	"fmt"
)

func main() {
	meuarray := [5]int{}

	meuarray[0] = 11
	meuarray[1] = 22
	meuarray[2] = 33
	meuarray[3] = 44
	meuarray[4] = 55

	for _, y := range meuarray {
		fmt.Println(y)
	}

	fmt.Printf("%T", meuarray)
}
