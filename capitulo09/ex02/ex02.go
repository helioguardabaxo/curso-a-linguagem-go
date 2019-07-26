package main

import (
	"fmt"
)

func main() {
	meuslice := []int{1, 2, 3, 4, 76, 99, 10, 12, 92}
	for _, y := range meuslice {
		fmt.Println(y)
	}

	fmt.Printf("%T", meuslice)
}
