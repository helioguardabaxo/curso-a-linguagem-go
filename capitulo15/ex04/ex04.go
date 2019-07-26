package main

import (
	"fmt"
	"sort"
)

func main() {
	cavaleiros := []string{"Pégasus", "Dragão", "Andrômeda", "Cisne", "Fênix"}
	idades := []int{12, 14, 21, 15, 11}

	sort.Ints(idades)
	fmt.Println(idades)

	sort.Strings(cavaleiros)
	fmt.Println(cavaleiros)
}
