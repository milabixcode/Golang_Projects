package main

import (
	"fmt"
)

func main() {

	var horaInicio, minutoInicio, horaFinal, minutoFinal int
	fmt.Scanln(&horaInicio, &minutoInicio, &horaFinal, &minutoFinal)
	tempoJogoHoras, tempoJogoMinutos := calcularTempo(horaInicio, minutoInicio, horaFinal, minutoFinal)
	fmt.Printf("O JOGO DUROU %d HORA(S) E %d MINUTO(S)\n", tempoJogoHoras, tempoJogoMinutos)
}

func calcularTempo(horaInicio, minutoInicio, horaFinal, minutoFinal int) (int, int) {

	var resultadoHoras, resultadoMinutos int

	if horaFinal > horaInicio {
		resultadoHoras = horaFinal - horaInicio
		resultadoMinutos = horaFinal*60 + 

		
	

		
	} else if horaFinal < horaInicio {
		resultadoHoras = 24 - (horaInicio - horaFinal)
	} else if horaFinal == horaInicio && minutoFinal == minutoInicio {
		resultadoHoras = 24
	}

	if minutoInicio >= 1 && minutoInicio < 60 && minutoFinal >= 1 && minutoFinal < 60 {
		resultadoMinutos = minutoFinal - minutoInicio
	}
	return resultadoHoras, resultadoMinutos
}
