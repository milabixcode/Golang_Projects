package main

import "fmt"

func generica(interf interface{}){
	fmt.Println(interf)
}

func main() {
	generica(1)
	generica("String")
	generica(true)

	//Ex:Println recebe várias interfaces, ele recebe quantos parametros de que tipo eu quiser 
	fmt.Println(1, "String", true)
}
