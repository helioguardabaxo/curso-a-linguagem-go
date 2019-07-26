package main

import (
	"fmt"
)

func main() {
	slice := [][]string{
		[]string{
			"joão",
			"trabalhador",
			"prefeitop",
		},
		[]string{
			"lula",
			"lulinha",
			"ladrão",
		},
		[]string{
			"picolé",
			"de chuchu",
			"sem cargo",
		},
	}

	for _, y := range slice {
		fmt.Println(y)
	}

	for _, y := range slice {
		fmt.Println(slice[0])
		for _, x := range y {
			fmt.Println("\t", x)
		}
	}
}
