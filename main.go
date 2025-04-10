package main

import (
	"math/rand"
)

func main() {

	votos := []int{}
	for i := 0; i < 10; i++ {
		voto := rand.Intn(5) + 1
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
