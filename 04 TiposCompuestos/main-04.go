package main

import "fmt"

func main() {
	// arrays()
	// slice1()
	// slice2()
	// slice3()
	// slicing()
	// sliceMake()
	// sliceMultiD()
	mapas1()
}

func arrays() {
	arr := [5]int{2, -5, 8, 3, -2}

	fmt.Printf("Arreglo de tipo %T\n", arr)
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func slice1() {
	miSlice := []int{3, 8, -1, 4, -9, 15, 68, -42, 2, -59}

	fmt.Printf("\nSlice de tipo %T\n", miSlice)
	for i, v := range miSlice {
		fmt.Println(i, v)
	}
}

func slice2() {
	miSlice := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(miSlice[:5])
	fmt.Println(miSlice[5:])
	fmt.Println(miSlice[2:7])
	fmt.Println(miSlice[1:6])
}

func slice3() {
	miSlice := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println("\n", miSlice)

	miSlice = append(miSlice, 52)
	fmt.Println(miSlice)

	miSlice = append(miSlice, 53, 54, 55)
	fmt.Println(miSlice)

	y := []int{56, 57, 58, 59, 60}
	miSlice = append(miSlice, y...)
	fmt.Println(miSlice)
}

func slicing() {
	miSlice := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println("\n", miSlice)
	y := append(miSlice[:3], miSlice[6:]...)
	fmt.Println(y)
}

func sliceMake() {
	sliceMx := make([]string, 32, 32)
	sliceMx = []string{"Aguascalientes", "Baja California", "Baja California Sur", "Campeche",
		"Coahuila de Zaragoza", "Colima", "Chiapas", "Chihuahua", "Distrito Federal", "Durango",
		"Guanajuato", "Guerrero", "Hidalgo", "Jalisco", "México", "Michoacán de Ocampo", "Morelos",
		"Nayarit", "Nuevo León", "Oaxaca", "Puebla", "Querétaro", "Quintana Roo", "San Luis Potosí",
		"Sinaloa", "Sonora", "Tabasco", "Tamaulipas", "Tlaxcala", "Veracruz de Ignacio de la Llave",
		"Yucatán", "Zacatecas"}

	nSlice := len(sliceMx)
	fmt.Println("\n", "La longitud es", len(sliceMx))
	fmt.Println("La capacidad es", cap(sliceMx))

	for i := 0; i < nSlice; i++ {
		fmt.Println(i+1, sliceMx[i])
	}
}

func sliceMultiD() {
	sliceX := []string{"batman", "jefe", "vestido negro"}
	sliceY := []string{"robin", "ayudante", "vestido de colores"}

	sliceM := [][]string{sliceX, sliceY}

	for i, v := range sliceM {
		fmt.Println("\nRegistro", i+1)
		for _, v2 := range v {
			fmt.Print("\t", v2)
		}
	}
}

func mapas1() {
	mapa1 := map[string][]string{
		"eduar_tua":    []string{"computadoras", "playa", "montaña"},
		"carlos_ramon": []string{"leer", "comprar", "musica"},
		"juan_bimba":   []string{"helado", "pintar", "bailar"}}

	for key1, val1 := range mapa1 {
		fmt.Println("\nRegistro:", key1)
		for _, val2 := range val1 {
			fmt.Print("\t", val2)
		}
	}
	println("")
	mapa1["angel_barrera"] = []string{"ultimate", "comer", "programar"}
	for key1, val1 := range mapa1 {
		fmt.Println("\nRegistro:", key1)
		for _, val2 := range val1 {
			fmt.Print("\t", val2)
		}
	}
	println("")
	if _, ok := mapa1["juan_bimba"]; ok {
		delete(mapa1, "juan_bimba")
	}
	for key1, val1 := range mapa1 {
		fmt.Println("\nRegistro:", key1)
		for _, val2 := range val1 {
			fmt.Print("\t", val2)
		}
	}
}
