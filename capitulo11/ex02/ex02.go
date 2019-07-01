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

	meumapa := make(map[string]pessoa)

	meumapa["Jorge"] = pessoa{
		nome:             "Jorge",
		sobrenome:        "Guedes",
		saboresFavoritos: []string{"Pistache", "Abacaxi"},
	}

	meumapa["Paulo"] = pessoa{
		nome:             "Paulo",
		sobrenome:        "Oliveira",
		saboresFavoritos: []string{"Morango", "Chocolate"},
	}

	for _, y := range meumapa {
		fmt.Printf("Meu nome é %s e meus sabores favoritos são: \n", y.nome)
		for _, z := range y.saboresFavoritos {
			fmt.Println(" - ", z)
		}

	}

}
