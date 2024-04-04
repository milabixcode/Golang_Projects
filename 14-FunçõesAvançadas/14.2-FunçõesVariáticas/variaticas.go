package main

import "fmt"

//Função recebe de 1 a n ints
func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}


func main() {
	totalDaSoma := soma(1,2,3,4,5)
	fmt.Println(totalDaSoma)

}