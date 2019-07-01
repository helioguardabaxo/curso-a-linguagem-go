package main

import (
	"fmt"
)

func retornaOutraFunc() func() {
	return func() {
		fmt.Println("Essa função foi chamada por outra função!")
	}
}

func main() {
	a := retornaOutraFunc()
	a()
}
