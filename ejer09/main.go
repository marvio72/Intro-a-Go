package main

import "fmt"

func main() {
	paises := make(map[string]string)
	fmt.Println(paises)

	paises["México"] = "CMX"
	paises["Argentina"] = "Buenos Aires"

	fmt.Println(paises)

	campeonato := map[string]int{
		"Barcelona":   39,
		"Chivas":      23,
		"Real Madrid": 11,
		"Boca":        32}

	fmt.Println(campeonato)

	campeonato["River Plate"] = 32
	delete(campeonato, "Real Madrid")
	fmt.Println(campeonato)

	for equipo, puntaje := range campeonato {
		fmt.Printf("El equipo %s, tiene un puntaje de: %d \n", equipo, puntaje)
	}
	puntaje, existe := campeonato["Chivas"]
	fmt.Printf("El puntaje capturado es %d, y el equipo existe %t \n", puntaje, existe)
}
