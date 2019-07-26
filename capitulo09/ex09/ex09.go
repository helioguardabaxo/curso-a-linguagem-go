package main

import (
	"fmt"
)

func main() {
	servidores := map[string][]string{
		"DaSilva_Dino":   []string{"futebol", "comer"},
		"Simpson_Hommer": []string{"baseball", "comer", "arrumar encrenca"},
		"Mouse_Mickey":   []string{"basquete", "queijo"},
	}

	servidores["Mamlsteen_Yngwie"] = []string{"tocar guitarra creme", "msotrar a lingua"}

	for x, y := range servidores {
		fmt.Println(x)
		for i, j := range y {
			fmt.Println("\t", i, " - ", j)
		}
	}
}
