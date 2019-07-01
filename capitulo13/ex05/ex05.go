package main

import (
	"fmt"
	"math"
)

type quadrado struct {
	lado float64
}

type circulo struct {
	raio float64
}

type info interface {
	area()
}

func medida(i info) {
	i.area()
}

func (q quadrado) area() {
	area := q.lado * q.lado
	fmt.Println("Área do quadrado:", area)
}

func (c circulo) area() {
	area := 2 * math.Pi * c.raio
	fmt.Println("Área do círculo:", area)
}

func main() {
	quadrado1 := quadrado{
		lado: 15.0,
	}

	circulo1 := circulo{
		raio: 5.25,
	}

	quadrado1.area()
	circulo1.area()

	//Usando interface
	medida(quadrado1)
	medida(circulo1)
}
