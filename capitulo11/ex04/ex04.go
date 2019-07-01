package main

import (
	"fmt"
)

type cliente struct {
	saboresFavoritos     []string
	sorveteriasFavoritas map[string]string
}

func main() {
	cliente1 := cliente{
		saboresFavoritos: []string{"Pistache", "Ovomaltine"},
		sorveteriasFavoritas: map[string]string{
			"Rio Pardo":          "Palácio do Sorvete",
			"São Caetano do Sul": "Lips",
			"Santo André":        "Mcdonalds",
		},
	}

	fmt.Println(cliente1)

	fmt.Println("Sabores favoritos:")
	for _, y := range cliente1.saboresFavoritos {
		fmt.Println(" - ", y)
	}
	fmt.Println("Sorveterias favoritas: ")
	for _, z := range cliente1.sorveteriasFavoritas {
		fmt.Println(" - ", z)
	}

}
