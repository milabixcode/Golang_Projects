package main

import "fmt"

func main(){
	fmt.Println("If e Else")

	numero := 10

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else if numero == 15 {
		fmt.Println("Número é igual a 15")
	} else {
		println("É menor que 15")
	}

	//if init
	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Número é maior que zero")
	} else {
		fmt.Println("Número é menor que zero")
	}
	// Eu não consigo acessar a variavel outroNumero fora do escopo
	//do if, variavel fica limitada
	//fmt.Println(outroNumero)
}