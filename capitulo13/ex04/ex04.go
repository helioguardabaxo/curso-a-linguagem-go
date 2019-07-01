package main

import (
	"fmt"
)

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

func main() {
	pessoa1 := pessoa{
		nome:      "Jose",
		sobrenome: "Pereira",
		idade:     45,
	}

	pessoa1.exibePessoa()
}

func (p pessoa) exibePessoa() {
	fmt.Println("O nome completo dessa pessoa é", p.nome, p.sobrenome, ".")
	fmt.Println(p.nome, p.sobrenome, "tem", p.idade, "anos.")
}
