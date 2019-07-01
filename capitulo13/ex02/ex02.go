package main

import (
	"fmt"
)

func main() {
	si := []int{19, 20, 21}
	si2 := []int{12, 13, 14}
	fmt.Println(funcao1(si...))
	fmt.Println(funcao2(si2))
}

func funcao1(x ...int) int {
	total := 0
	for _, j := range x {
		total += j
	}
	return total
}

func funcao2(x []int) int {
	total := 0
	for _, j := range x {
		total += j
	}
	return total
}
