package main

import (
	"fmt"
	"math"
)

type human interface {
	presentar()
}
type persona struct {
	nombre, apellido string
	edad             int
}
type agenteSecetro struct {
	persona
	lpm bool
}

func (p persona) presentar() {
	fmt.Println("Hola, me llamo", p.nombre, p.apellido)
	fmt.Println("Y tengo", p.edad)
}
func (as agenteSecetro) presentar() {
	fmt.Println("Hola, soy el agente", as.nombre, as.apellido)
}
func bar(h human) {
	switch h.(type) {
	case persona:
		fmt.Println("La persona de nombre ", h.(persona).nombre, "saluda desde bar")
	case agenteSecetro:
		fmt.Println("El agente", h.(agenteSecetro).nombre, "esta en bar")
	}
	fmt.Println(h, "es argumento de bar")
}

type circulo struct {
	radio float32
}
type rectangulo struct {
	lado, altura float32
}

func (c circulo) area() float32 {
	return float32(math.Pi * math.Pow(float64(c.radio), 2.0))
}
func (r rectangulo) area() float32 {
	return r.altura * r.lado
}

type forma interface {
	area() float32
}

func info(f forma) {
	fmt.Println("El area es ", f.area())
}

func foo() (int, string) {
	return 4, "mi numero es "
}

func foo2() {
	defer func() {
		fmt.Println("funcion anonima en foo")
	}()
	fmt.Println("saludos desde funcion foo")
}

func sumaElems(elems ...int) int {
	total := 0
	for _, val := range elems {
		total += val
	}
	return total
}

var gg func() string = func() string {
	return "Checate este mensaje"
}

func clousure() func() int {
	fmt.Println("Hola desde clousure")
	x := 0
	return func() int {
		fmt.Println("Hola desde hijo de clousure")
		x++
		return x
	}
}
func callback(call1 func([]int) int, items []int) int {
	return call1(items) + 1
}
func sumaSlice(coleccion []int) int {
	if n := len(coleccion); n == 0 {
		return 0
	} else if n == 1 {
		return coleccion[0]
	} else {
		return coleccion[0] + coleccion[n-1]
	}
}

func main() {
	defer foo2()
	p1 := persona{
		nombre:   "Angel",
		apellido: "Barrera",
	}
	as1 := agenteSecetro{
		persona: persona{
			nombre:   "James",
			apellido: "Bond",
		},
		lpm: true,
	}

	p1.presentar()
	as1.presentar()
	bar(p1)
	bar(as1)

	numero, msj := foo()
	fmt.Println(msj, numero)

	items := []int{5, 8, -8, 7, 2, -6, 4}
	fmt.Println("La suma de elementos es", sumaElems(items...))

	circ := circulo{
		radio: 3,
	}
	rec := rectangulo{
		altura: 2,
		lado:   2,
	}
	info(circ)
	info(rec)

	fmt.Printf("La variable gg es de tipo %T\n", gg)
	fmt.Println(gg())

	clous := clousure()
	fmt.Println(clous())
	fmt.Println(clous())
	fmt.Println(clous())

	fmt.Println("La suma de elementos en el callback es", callback(sumaSlice, items))

}
