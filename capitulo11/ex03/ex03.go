package main

import (
	"fmt"
)

type veiculo struct {
	portas int
	cor    string
}

type caminhonete struct {
	veiculo
	tracaoNasQuatro bool
}

type sedan struct {
	veiculo
	modeloLuxo bool
}

func main() {
	caminhonete1 := caminhonete{
		veiculo: veiculo{
			portas: 4,
			cor:    "Azul",
		},
		tracaoNasQuatro: true,
	}

	sedan1 := sedan{
		veiculo: veiculo{
			portas: 4,
			cor:    "Preta",
		},
		modeloLuxo: true,
	}

	fmt.Println(caminhonete1)
	fmt.Println(sedan1)

	fmt.Println(caminhonete1.portas)
	fmt.Println(sedan1.cor)
}
