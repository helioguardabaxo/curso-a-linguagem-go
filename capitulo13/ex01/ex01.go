package main

import (
	"fmt"
)

func main() {
	fmt.Println(retornaUmInt())
	fmt.Println(retornaUmIntUmaString())
}

func retornaUmInt() int {
	return 10
}

func retornaUmIntUmaString() (int, string) {
	return 10, "Dez"
}
