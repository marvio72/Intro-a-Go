package main

import (
	"fmt"
)

// Calculo : Variable que contiene función anonima.
var Calculo func(int, int) int = func(num1 int, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Printf("Suma 5 + 8 = %d \n", Calculo(5, 8))

	// Restamos
	Calculo = func(num1 int, num2 int) int {
		return num1 - num2
	}
	fmt.Printf("Restamos 6 - 4 = %d \n", Calculo(6, 4))

	// Dividimos
	Calculo = func(num1 int, num2 int) int {
		return num1 / num2
	}
	fmt.Printf("Dividimos 24 / 8 = %d \n", Calculo(24, 8))

	// Multiplicamos
	Calculo = func(num1 int, num2 int) int {
		return num1 * num2
	}
	fmt.Printf("Multiplicamos 11 * 6 = %d \n", Calculo(11, 6))

	Operaciones()

	/* CLOSURES */
	tablaDel := 5
	MiTabla := Tabla(tablaDel)
	for i := 1; i < 11; i++ {
		fmt.Println(MiTabla())
	}

}

// Operaciones : Otra manera de crear una función anónima
func Operaciones() {
	resultado := func() int {
		var a int = 23
		var b int = 13
		return a + b
	}
	fmt.Println(resultado())

}

// Tabla : Funcion tabla de multiplicar
func Tabla(valor int) func() int {
	numero := valor
	secuencia := 0
	return func() int {
		secuencia++
		return numero * secuencia
	}
}
