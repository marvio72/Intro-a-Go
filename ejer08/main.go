package main

import "fmt"

// var tabla [10]int

func main() {
	// tabla[0] = 1
	// tabla[5] = 15

	//otra manera de asignar un arreglo
	tabla := [10]int{4, 2, 4, 65, 43, 23, 65, 32, 98, 102}

	// para presentar los datos de la tabla en forma vertical
	for i := 0; i < len(tabla); i++ {

		fmt.Println(tabla[i])
	}

}
