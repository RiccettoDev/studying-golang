// Crie um tipo "quadrado"
// Crie um tipo "círculo"
// Crie um método "área" para cada tipo que calcule e retorne a área da figura
// Área do círculo: 2 * π * raio
// Área do quadrado: lado * lado
// Crie um tipo "figura" que defina como interface qualquer tipo que tiver o método "área"
// Crie uma função "info" que receba um tipo "figura" e retorne a área da figura
// Crie um valor de tipo "quadrado"
// Crie um valor de tipo "círculo"
// Use a função "info" para demonstrar a área do "quadrado"
// Use a função "info" para demonstrar a área do "círculo"

package main

import (
	"fmt"
)

type quadrado struct {
	ladoA float64
	ladoB float64
}

type circulo struct {
	raio float64
}

func (q quadrado) area() float64 {
	result := q.ladoA * q.ladoB
	return result
}

func (c circulo) area() float64 {
	result := 2 * 3.14159 * c.raio
	return result
}

type figura interface {
	area() float64
}

func info(f figura) float64 {
	return f.area()
}

func main() {
	quadrado1 := quadrado{
		ladoA: 2,
		ladoB: 2,
	}

	circulo1 := circulo{
		raio: 5.25,
	}

	fmt.Println(info(quadrado1))
	fmt.Println(info(circulo1))
}
