package main

import (
	"fmt"
)

func main() {
	func() {
		fmt.Println("Esta é uma função anônima!")
	}()
}
