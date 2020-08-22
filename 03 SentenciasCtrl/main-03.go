package main

import "fmt"

func main() {
	// forAnidado()
	// fmt.Print("\n")
	// forBreakContinue()
	// fmt.Print("\n")
	// ifElse()
	// fmt.Print("\n")
	// switchBool()
	// fmt.Print("\n")
	// switchString()
	// fmt.Print("\n")
	ejercicios()
}

func forAnidado() {
	for i := 0; i < 5; i++ {
		fmt.Println("iteracion", i)
		for j := 0; j < 3; j++ {
			fmt.Println("\titeracion", j)
		}
	}
}

func forBreakContinue() {
	x := 1
	for {
		x++
		if x > 12 {
			break
		}
		if x%2 != 0 {
			continue
		}
		fmt.Println(x)
	}
}

func ifElse() {
	if a := 4; a > 6 {
		fmt.Println("Cumplio el if")
	} else if a < 0 {
		fmt.Println("a es negativo")
	} else if a == 0 {
		fmt.Println("a es 0")
	} else {
		fmt.Println("se fue al default, a vale", a)
	}
}

func switchBool() {
	a := 7
	switch {
	case a%2 == 0, a > 0, a != 1:
		println("a es par, mayor a cero o diferente de uno")
		fallthrough
	case a < 12:
		println("Entro al 2do caso")
	default:
		println("Entro al caso por defecto")
	}
}

func switchString() {
	fruto := "cereza"
	switch fruto {
	case "fresa", "cereza", "mora":
		println("frutos rojos")
	case "uva", "manzana", "pera":
		println("frutos verdes")
	default:
		println("no hay un color definido para ese fruto")
	}
}

func ejercicios() {
	for i := 65; i < 91; i++ {
		fmt.Printf("\n%d\t", i)
		for j := 0; j < 2; j++ {
			fmt.Printf("%#U\t", i)
		}
	}
	fmt.Print("\n\n")

	anioNac := 1996
	for anioNac < 2020 {
		fmt.Print(anioNac, "\t")
		anioNac++
	}
	fmt.Print("\n\n")

	for indice := 1; indice <= 20; indice++ {
		fmt.Printf("El resto de %d/4 es %v\n", indice, indice%4)
	}
}
