package main

import (
	"fmt"
)

func main() {
	a := func() {
		fmt.Println("Oi, eu sou um função atribuída a uma variável")
	}
	a()

}
