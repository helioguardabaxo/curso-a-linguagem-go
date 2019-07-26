package main

import (
	"fmt"
)

type pessoa struct {
	nome  string
	idade int
}

func mudeMe(p *pessoa) {
	(*p).nome = "Jose"
	(*p).idade = 88
}

func main() {
	pessoa1 := pessoa{
		nome:  "Hélio",
		idade: 32,
	}

	fmt.Println(pessoa1)

	mudeMe(&pessoa1)
	fmt.Println(pessoa1)
}
