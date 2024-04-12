package main

import "fmt"

//Canais com Buffer
//Especifica capacidade para o canal
//Operação de receber e enviar dados são operações bloqueantes
//
func main() {
	canal := make(chan string, 200)
	canal <- "Olá Mundo!"
	canal <- "Programando em Go"
	canal <- "Programando em Go de novo"

	mensagem :=  <-canal
	mensagem2 := <- canal
	mensagem3 := <- canal

	fmt.Println(mensagem)
	fmt.Println(mensagem2)
	fmt.Println(mensagem3)
}