package main

import (
	"fmt"
)

type pessoa struct {
	nome             string
	sobrenome        string
	saboresFavoritos []string
}

func main() {

	pessoa1 := pessoa{
		nome:             "Jorge",
		sobrenome:        "Guedes",
		saboresFavoritos: []string{"Pistache", "Abacaxi"},
	}

	pessoa2 := pessoa{
		nome:             "Paulo",
		sobrenome:        "Oliveira",
		saboresFavoritos: []string{"Morango", "Chocolate"},
	}

	fmt.Println(pessoa1.nome, pessoa1.sobrenome)
	fmt.Println("Sabores favoritos:")
	for x, y := range pessoa1.saboresFavoritos {

		fmt.Printf("%d - %s\n", x, y)
	}

	fmt.Println(pessoa2.nome, pessoa2.sobrenome)
	fmt.Println("Sabores favoritos:")
	for x, y := range pessoa2.saboresFavoritos {

		fmt.Printf("%d - %s\n", x, y)
	}

}
