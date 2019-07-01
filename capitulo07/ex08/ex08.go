package main

import (
	"fmt"
)

func main() {
	nota := 57

	switch {
	case nota < 60:
		fmt.Println("Reprovado")
	case nota > 60 && nota < 80:
		fmt.Println("Aprovado")
	}
}
