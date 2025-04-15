package main

import (
	"fmt"
)

func main() {

	votos := []int{}
	for i := 0; i < 10; i++ {
		var voto int
		println("Introduce el voto (1-5):")
		_, err := fmt.Scan(&voto)
		if err != nil || voto < 1 || voto > 5 {
			println("Voto no válido, introduce un número entre 1 y 5")
			i--
			continue
		}
		votos = append(votos, voto)
	}

	votosContados := make(map[int]int)
	for _, voto := range votos {
		votosContados[voto]++
	}
	for i := 1; i <= 5; i++ {
		if votosContados[i] > 0 {
			println("Voto", i, ":", votosContados[i])
		}
	}
	if votosContados[4]+votosContados[5] > votosContados[1]+votosContados[2] {
		println("Buen resultado")
	} else {
		println("Resultado mejorable")
	}

}
