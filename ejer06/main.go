package main

import (
	"fmt"
)

func main() {
	// fmt.Println(uno(5))

	// numero, estado := dos(2)

	// fmt.Println(numero)
	// fmt.Println(estado)

	fmt.Println(calculo(5, 46))
	fmt.Println(calculo(2, 46, 7, 23, 38))
	fmt.Println(calculo(7))
	fmt.Println(calculo(5, 46, 23, 43, 456, 17, 29, 323))

}

func uno(numero int) (z int) {
	z = numero * 2
	return
}

func dos(numero int) (int, bool) {
	if numero == 1 {
		return 5, false
	} else {
		return 10, true
	}
}

// Parametros variables
func calculo(numero ...int) int {
	total := 0
	for item, num := range numero {
		total = total + num
		fmt.Printf("\n Item %d \n", item)
	}
	return total
}
