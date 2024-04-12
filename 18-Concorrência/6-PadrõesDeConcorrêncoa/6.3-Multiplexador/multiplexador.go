package main


import (
	"fmt"
	"time"
	"math/rand"
)

//Mutiplexador: pegar 2 ou mais canais e juntar em 1 só

func main () {

	canal := multiplexar(escrever("Olá, Mundo"), escrever("Programando em Go!()"))

	for i := 0; i < 190; i++ {
		fmt.Println(<-canal)
	}

}

func multiplexar(canalDeEntrada1, canalDeEntrada2 <- chan string) <- chan string {
	canalDeSaida := make(chan string)

	//Chamada para um GoRoutine anônima com for infinito
	//Select: vai ver qual dos dois canais tem msg para ser lida e joga para o canal de saída
	go func() {
		for {
			select {
			case mensagem := <- canalDeEntrada1:
				canalDeSaida <- mensagem
			case mensagem := <- canalDeEntrada2:
				canalDeSaida <- mensagem
			}
		}
	} ()

	return canalDeSaida
}

func escrever(texto string) <- chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", texto)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()

	return canal
}