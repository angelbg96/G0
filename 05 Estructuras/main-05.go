package main

import "fmt"

type persona struct {
	nombre, apellido string
	helados          []string
}
type vehiculo struct {
	puertas int
	color   string
}
type camion struct {
	vehiculo
	cuatroRuedas bool
}
type sedan struct {
	vehiculo
	lujoso bool
}

func main() {
	// estructura1()
	// mapa1()
	// estructuraEmb()
	estruturaAnon()
}

func estructura1() {
	p1 := persona{
		nombre:   "Angel",
		apellido: "Barrera",
		helados:  []string{"limon", "mango", "grosella"},
	}

	p2 := persona{
		nombre:   "Pepito",
		apellido: "Perez",
		helados:  []string{"chocolate", "vainilla", "fresa"},
	}

	fmt.Println("\n", p1.nombre, p1.apellido)
	fmt.Print("Helados fav :")
	for _, v := range p1.helados {
		fmt.Print("\t", v)
	}
	fmt.Println("\n", p2.nombre, p2.apellido)
	fmt.Print("Helados fav :")
	for _, v := range p2.helados {
		fmt.Print("\t", v)
	}
}

func mapa1() {
	m1 := map[string]persona{
		"barrera_gal": persona{
			nombre:   "Angel",
			apellido: "Barrera",
			helados:  []string{"limon", "mango", "grosella"},
		},
		"perez": persona{
			nombre:   "Pepito",
			apellido: "Perez",
			helados:  []string{"chocolate", "vainilla", "fresa"},
		}}

	for key, val := range m1 {
		fmt.Println("\nRegistro", key)
		fmt.Println("\t", val.nombre, val.apellido)
		fmt.Print("\tHelados fav :")
		for _, helado := range val.helados {
			fmt.Print("\t", helado)
		}
	}
}

func estructuraEmb() {
	ca1 := camion{
		vehiculo: vehiculo{
			puertas: 2,
			color:   "blanco",
		},
		cuatroRuedas: true,
	}
	se1 := sedan{
		vehiculo: vehiculo{
			puertas: 4,
			color:   "rojo",
		},
		lujoso: false,
	}

	println("\n", ca1.color, ca1.puertas, ca1.cuatroRuedas)
	println(se1.color, se1.puertas, se1.lujoso)
}

func estruturaAnon() {
	miEstruc := struct {
		nombre      string
		edad        int16
		altura      float32
		tecnologias map[string][]string
	}{
		nombre: "Angel",
		edad:   23,
		altura: 1.68,
		tecnologias: map[string][]string{
			"fronted": []string{"html", "css", "js", "vueJS"},
			"backend": []string{"php", "java", "go"}},
	}
	fmt.Println("\n", miEstruc.nombre, miEstruc.edad, "años", miEstruc.altura, "m")
	for key, val := range miEstruc.tecnologias {
		fmt.Print("\t", key, " :")
		for _, val2 := range val {
			fmt.Print("\t", val2)
		}
		fmt.Println("")
	}
}
