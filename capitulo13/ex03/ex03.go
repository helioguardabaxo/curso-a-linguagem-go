package main

import (
	"fmt"
)

func main() {
	fmt.Println("Primeira linha")
	defer fmt.Println("Segunda linha")
	fmt.Println("Terceira linha")
}
