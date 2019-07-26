package main

import (
	"fmt"
)

func main() {
	meuslice := []int{1, 2, 3, 44, 7, 90, 12, 34}

	fmt.Println(meuslice[:3])
	fmt.Println(meuslice[4:])
	fmt.Println(meuslice[1:7])
	fmt.Println(meuslice[2 : len(meuslice)-1])
}
