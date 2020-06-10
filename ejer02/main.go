package main

import (
	"fmt"
	"strconv"
)

// declaración de variables globales
var numero int
var texto string
var status bool = true

func main() {
	numero2, numero3, numero4, texto, status := 2, 23, 642, "Este es mi texto", false

	var moneda int64 = 77

	numero2 = int(moneda)

	texto = strconv.Itoa(int(moneda))

	fmt.Println(numero2)
	fmt.Println(numero3)
	fmt.Println(numero4)
	fmt.Println(texto)
	fmt.Println(status)

	mostrarStatus()
}

func mostrarStatus() {
	fmt.Println(status)
}
