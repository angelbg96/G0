package main

import "fmt"

var x int
var y string
var z bool

type (
	oso        string
	autorizado bool
	numero     int
)

var panda oso
var nuevoUsuario autorizado

func main() {
	miVar1, miVar2, miVar3 := 9, 3.2, true
	stringConca := fmt.Sprintf("los valores de mis 3 variables son:\t%v\t%v\t%v\n", miVar1, miVar2, miVar3)
	var num1 numero

	fmt.Printf(" x es de tipo %T y vale %d por default\n", x, x)
	fmt.Printf(" y es de tipo %T y vale %s por default\n", y, y)
	fmt.Printf(" z es de tipo %T y vale %t por default\n", z, z)

	fmt.Println("\n", stringConca)

	panda = "panda"
	fmt.Printf("soy un %s de tipo %T\n", panda, panda)
	fmt.Printf("Acceso %t para el nuevo usuario\n\n", nuevoUsuario)

	fmt.Printf("Soy un tipo %T que realmente es %T\n", num1, int(num1))
}
