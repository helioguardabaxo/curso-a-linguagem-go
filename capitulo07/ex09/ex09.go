package main

import (
	"fmt"
)

func main() {
	esporteFavorito := "Basquete"
	switch esporteFavorito {
	case "Futebol":
		fmt.Println("Futebol não é tão legal assim.")
	case "Basquete":
		fmt.Println("Aí sim! Vamos assistir o papai Lebron?!")
	}
}
