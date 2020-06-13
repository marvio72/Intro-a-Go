package main

import (
	"fmt"
)

func main() {
	i := 1
	for i < 10 {
		fmt.Println(i)
		i++
	}

	fmt.Println("---------------------------------------")
	//Otra manera mas parecida al for de javaScript
	for j := 1; j <= 10; j++ {
		fmt.Println(j)
	}

	// Rompiendo loops
	fmt.Println("---------------------------------------")
	numero := 0
	for {
		fmt.Println("Continuo")
		fmt.Println("ingrese el número secreto")
		fmt.Scanln(&numero)
		if numero == 100 {
			break
		}

	}

	fmt.Println("---------------------------------------")
	var k = 0
	for k < 10 {
		fmt.Printf("\n Valor de k: %d", k)
		if k == 5 {
			fmt.Printf(" multiplicamos por 2 \n")
			k = k * 2
			continue
		}
		fmt.Printf("    Pasó por aquí \n")
		k++
	}

	fmt.Println("---------------------------------------")
	var l int = 0

RUTINA:
	for l < 10 {
		if l == 4 {
			l = l + 2
			fmt.Println("voy a RUTINA sumando 2 a l")
			goto RUTINA
		}
		fmt.Printf("Valor de l: %d\n", l)
		l++
	}
}
