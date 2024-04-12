package main

import "fmt"
import "time"


//deadlock: não tem mais nenhum lugar que ta enviiando dado para o canal mas o canal está esperando receber dado
//não é identificado em compilação, só na execução 
func main(){
	canal := make(chan string)

	go escrever("Olá mundo!", canal)

	fmt.Println("Depois da função escrever começar a ser executada")

	for mensagem := range canal {
		fmt.Println(mensagem)
	}
	fmt.Println("Fim do Programa!")
}

		
func escrever(texto string, canal chan string) {
	for i:=0; i < 5; i++ {
		canal <- texto
		time.Sleep(time.Second)
	}

	close(canal)
}