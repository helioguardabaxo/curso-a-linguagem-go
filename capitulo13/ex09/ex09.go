package main

import (
	"fmt"
)

func retornaArgumento() {
	fmt.Println("Essa função foi chamada por outra função!")
}

func recebeOutraFuncao(f func()) {
	fmt.Println("Preste a atenção:")
	f()
}

func main() {
	recebeOutraFuncao(retornaArgumento)
}
