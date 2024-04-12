package main

import "fmt"

//Fila de tarefas para serem executadas e tem vários workers (processos) pegando itens dessa fila de maneira independente
//Mais deum processo rodando independente
func main () {
	//canal com buffer
	//sequ^ncia de numeros
	tarefas := make(chan int, 45)
	//armazena resultdos a medida oque são calculados
	resultados := make(chan int, 45)

	//Concorrência: 4 processos ao mesmo tempo: torna mais rápido
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)

	for i := 0; i < 45; i++ {
		tarefas <- i
	}
	close(tarefas)

	for i := 0; i < 45; i++ {
		resultado := <- resultados
		fmt.Println(resultado)
	}


}

func worker (tarefas <- chan int, resultados chan <- int) {
	for numero := range tarefas {
		resultados <- fibonacci(numero)
	}
}

func fibonacci(posicao int) int {
	if posicao <= 1 {
		return posicao
	}
	return fibonacci(posicao - 2) + fibonacci(posicao - 1)
}