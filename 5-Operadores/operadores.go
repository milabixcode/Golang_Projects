package main

import "fmt"

func main(){
	// ARITMETICOS
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 5 
	restoDivisao :=10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDivisao)

	//se tiver tipos diferentes dá erro
	var numero1 int16 = 10
	var numero2 int32 = 25
	soma2 := numero1 + int16(numero2)
	fmt.Println(soma2)
	
	// FIM ARITMÉTICOS
	
	//ATRIBUIÇÃO
	var variavel1 string = "String"
	variavel2 := "String2"
	fmt.Println(variavel1, variavel2)

	//FIM DOS OPERADORES DE ATRIBUIÇÃO

	//OPERADORES RELACIONAIS
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 < 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 != 2)

	//FIM DOS OPERADORES RELACIONAIS

	//OPERADORES LÓGICOS
	fmt.Println("---------------------")

	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)
	
	//FIM OPERADORES LÓGICOS

	//OPERADORES UNÁRIOS
	numero := 10
	numero ++
	numero += 15 //numero = numero + 15
	fmt.Println((numero))

	numero --
	numero -= 20

	fmt.Println(numero)

	numero *= 3
	numero /= 10
	numero %= 3

	fmt.Println(numero)

	//No Go não existe ++numero ou --numero
	//FIM OPERADORES UNÁRIOS

	//OPERADORES TERNÁRIOS
	//Não existe isso em Go
	// seria assim se existisse: texto := numero > 5 ? "Maior que 5" : "Menor que 5"
	//Refatorando:
	var texto string
	if numero > 5 {
		texto = "maior que 5"
	} else {
		texto = "menor que 5"
	}
	fmt.Println(texto)

}
