package main

import (
	"fmt"
)

func main() {
	anoNascimento := 1986
	for {
		if anoNascimento <= 2018 {
			fmt.Println(anoNascimento)
			anoNascimento++
		} else {
			break
		}

	}
}
