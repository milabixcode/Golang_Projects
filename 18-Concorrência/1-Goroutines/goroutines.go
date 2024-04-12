package main

import (
	"fmt"
	"time"
)

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}

// Concorrência != Paralelismo

// O Paralelismo acontece quando há duas ou mais tarefas que estão sendo executadas exatamente
// ao mesmo tempo e só é possível se tiver um processador de mais de núcleo (distribui tarefas 
//entre os núcleos)
// Tarefas que estão sendo executadas concorrente, não estão necessariamente executando ao mesmo tempo 
//(duas tarefas rodando e uma tarefa não esperaria a outra para rodar também - revezamento)

func main() {
	go escrever("Hello World!")
	escrever("Programando em Go!")
}
