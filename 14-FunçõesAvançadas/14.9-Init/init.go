package main

import "fmt"

var n int

func main() {
	fmt.Println("Função main sendo executada")
	fmt.Println(n)
}

func init() {
	fmt.Println("Função init sendo executada")
	n := 10
	fmt.Println(n)
}
