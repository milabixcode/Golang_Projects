package main

import "fmt"

func main(){
	var variavel1 string = "Variavel 1"
	variavel2 := "Variavel 2" //Inferência de tipo
	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var (
		variavel3 string = "Variavel 3"
		variavel4 string = "Variavel 4"
	)
	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "Variável 5", "Variável 6"
	fmt.Println(variavel5, variavel6)
}