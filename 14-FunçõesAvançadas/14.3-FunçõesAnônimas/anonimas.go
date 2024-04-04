package main

import "fmt"

func main() {

//Função que não precisa de nome
	retorno := func(texto string) string {
		return fmt.Sprintf("Recebido -> %s", texto)
	}("Passando Parâmetro")

		fmt.Println(retorno)
}
