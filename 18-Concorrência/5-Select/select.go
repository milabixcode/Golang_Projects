package main

import (
	"fmt"
	"time"
)

func main() {
	canal1, canal2 := make(chan string), make(chan string)

	//recebe a cada meio segundo
	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			canal1 <- "Canal 1"
		}
	}()

	//recebe a cada 2 segundos
	go func() {
		for {
			time.Sleep(time.Second * 2)
			canal2 <- "Canal 2"
		}
	}()

	//quando uma função demora mais que a outra para ser executada
	//select faz com que uma função não precise esperar a outra para executar
	for {
		select {
		case mensagemCanal1 := <- canal1:
			fmt.Println(mensagemCanal1)
		case mensagemCanal2 := <-canal2:
			fmt.Println(mensagemCanal2)
		}
	}
}
