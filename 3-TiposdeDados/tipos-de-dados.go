package main

import (
	"errors"
	"fmt"
)

func main() {
	var numero int = 1000000000000000000 //int sem especificaçao de bits vai de acordo com a arquitetura do computador
	numero1 := 20000000000000000
	fmt.Println(numero)
	fmt.Println(numero1)

	//uint = int sem sinal

	//alias
	// rune = int32 
	var numero3 rune = 123456
	fmt.Println(numero3)

	//byte = uint8
	var numero4 byte = 123
	fmt.Println(numero4)

	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 12300000000000.45
	fmt.Println(numeroReal2)

	//Inferência de Tipo
	numeroReal3 := 12345.67
	fmt.Println(numeroReal3)

////////////////////////////////////////////////////////////////////////////////////////////
//STRINGS
	//GO NÃO TEM CHAR

var str string = "Texto"
fmt.Println(str)

str2 := "Texto2"
fmt.Println(str2)

//Numero correspondente na tabela ASCII
char := 'B'
fmt.Println(char)

////////////////////////////////////////////////////////////////////////////////////////////!SECTION

// Todo tipo de dado tem seu valor inicial
// String = String vazio
// int = 0
// bool = false
// error = <nil>

var texto int16
fmt.Println(texto)

var booleano1 bool 
fmt.Println(booleano1)

var erro error
fmt.Println(erro)

//errors: pacote
var erro1 error = errors.New(("Erro interno"))
fmt.Println(erro1)





}