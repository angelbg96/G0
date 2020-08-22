package main

import "fmt"

const (
	_  = iota
	kb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
	tb = 1 << (iota * 10)
)

func main() {
	strings()
	ctes()
	ejercicios()
}

func strings() {
	s1 := "Hola GOlers"
	s2 := `Cadena tipo
		raw format`
	bs := []byte(s1)

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Printf("\nla cadena está codificada en %T\n", bs)
	for i, v := range bs {
		fmt.Printf("%d) %#X\t", i, v)
	}
	fmt.Printf("\ncon %q se imprime cualquier cadena\n\n\n", "%s")
}

func ctes() {
	fmt.Printf("%d\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t%b\n", gb, gb)
	fmt.Printf("%d\t\t%b\n", tb, tb)
}

func ejercicios() {
	num1 := 28
	num2 := 86
	fmt.Printf("El numero %d en binario es %b y en hexadecimal es %#X\n", num1, num1, num1)

	a := (num1 == num2) || (num1 > num2)
	b := (num1 != num2) || (num1 < num2)
	println("el valor de a es", a)
	println("el valor de b es", b)

	const (
		ct1         = 12.2
		ct2 float32 = 12.2
	)
	println("\ncte1 = ", ct1)
	println("cte2 = ", ct2)

	numSB := num1 >> 1
	fmt.Printf("\nEl numero %d en binario es %b y en hexadecimal es %#X\n\n", numSB, numSB, numSB)

	rawString := `"Al natural"...
	\no mas 'wacha'		este 			formato\`
	fmt.Println(rawString)

	const (
		anio0 = 2020 + iota
		anio1 = 2020 + iota
		anio2 = 2020 + iota
		anio3 = 2020 + iota
	)
	println("\n", anio0, anio1, anio2, anio3)
}
