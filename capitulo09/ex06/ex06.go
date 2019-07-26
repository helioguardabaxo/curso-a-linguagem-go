package main

import (
	"fmt"
)

func main() {
	slice := make([]string, 2, 2)
	slice = []string{"Acre", "Alagoas"}

	fmt.Println(slice)
	fmt.Println(len(slice), cap(slice))

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}
}
