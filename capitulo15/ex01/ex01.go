package main

import (
	"fmt"
)

func main() {
	x := 1000

	fmt.Printf("Valor de %v na memória: %v", x, &x)
}
