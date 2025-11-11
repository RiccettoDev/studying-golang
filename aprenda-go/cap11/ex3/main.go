// Crie um novo tipo: veículo
// O tipo subjacente deve ser struct
// Deve conter os campos: portas, cor
// Crie dois novos tipos: caminhonete e sedan
// Os tipos subjacentes devem ser struct
// Ambos devem conter "veículo" como struct embutido
// O tipo caminhonete deve conter um campo bool chamado "traçãoNasQuatro"
// O tipo sedan deve conter um campo bool chamado "modeloLuxo"
// Usando os structs veículo, caminhonete e sedan:
// Usando composite literal, crie um valor de tipo caminhonete e dê valores a seus campos
// Usando composite literal, crie um valor de tipo sedan e dê valores a seus campos
// Demonstre estes valores.
// Demonstre um único campo de cada um dos dois.

package main

import (
	"fmt"
)

type vehicle struct {
	doors int
	color string
}

type pickupTruck struct {
	vehicle
	tracaoNasQuatro bool
}

type sedan struct {
	vehicle
	modeloLuxo bool
}

func main() {

	saveiroVehicle := vehicle{
		doors: 2,
		color: "Black",
	}

	saveiro := pickupTruck{
		vehicle:         saveiroVehicle,
		tracaoNasQuatro: true,
	}

	hondaVehicle := vehicle{
		doors: 4,
		color: "Black",
	}

	honda := sedan{
		vehicle:    hondaVehicle,
		modeloLuxo: true,
	}

	x := ""
	if saveiro.tracaoNasQuatro {
		x = "com"
	} else {
		x = "sem"
	}
	fmt.Printf("Pickup Saveiro com %v portas, na cor %v %v tração nas quatro rodas\n", saveiro.doors, saveiro.color, x)

	y := ""
	if honda.modeloLuxo {
		y = "é um"
	} else {
		y = "não é um"
	}
	fmt.Printf("Sedan Honda com %v portas, na cor %v %v modelo de luxo\n", honda.doors, honda.color, y)

}
