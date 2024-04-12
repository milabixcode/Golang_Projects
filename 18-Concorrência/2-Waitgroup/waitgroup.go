package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

//Não quero uma função dependa da outra
//Quero sincronizar elas: quero que as duas tenham terminado antes do meu programa terminar de executar
//Waitgroup
	
	var waitGroup sync.WaitGroup
	//qtd de goRoutines que estarão rodando ao mesmo tempo
	waitGroup.Add(2)
	
	
	//função anônima
	//rodando concorrente
	go func () {
		escrever("Hello World")
		//qdo função terminar chama e tira uma função do "contador"
		waitGroup.Done() //(-1)
	} ()
		
	go func () {
			escrever("Wait Group")
			waitGroup.Done() //(-1)
	} ()

	// Fala para a função esperar a contagem das goRoutines chegar em 0
	waitGroup.Wait()
}	
		
func escrever(texto string) {
	for i:=0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}