package main

import "fmt"

var estado bool

func main() {
	estado = true
	if estado = false; estado == true {
		fmt.Println(estado)
	} else if estado == false {
		fmt.Println("Es Falso")
	}
}
