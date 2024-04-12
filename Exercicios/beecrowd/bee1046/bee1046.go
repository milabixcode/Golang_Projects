package main

import (
	"fmt"
)

func main() {

	var horaInicio, horaFinal int
	fmt.Scanln(&horaInicio, &horaFinal)
	tempoJogo := calcularTempoJogo(horaInicio, horaFinal)
	fmt.Printf("O JOGO DUROU %d HORA(S)\n", tempoJogo)
	
}

func calcularTempoJogo(horaInicio, horaFinal int) int {
	var resultado int
	if horaFinal > horaInicio {
		resultado = horaFinal - horaInicio
	} else if horaFinal < horaInicio {
		resultado = 24 - (horaInicio - horaFinal)
	} else if horaFinal == horaInicio {
		resultado = 24 
	}
	return resultado
}
