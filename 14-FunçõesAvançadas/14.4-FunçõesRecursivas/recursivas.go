package main

import "fmt"

//Função que retorna o número que está em uma determinada posição na SequÊncia de Fibonacci
//Fibonacci: o proximo numero é sempre a soma dos números anteriores
func fibonacci(posicao uint) uint {
	//preciso de condição de parada: overflow de pilha
	if posicao <= 1 {
		return posicao
	}
	return fibonacci(posicao - 2) + fibonacci(posicao - 1)
}


func main() {
	fmt. Println("Funções Recursivas")

	// 1 1 2 3 5 8 13

	posicao := uint(7)

	for i := uint(0); i < posicao; i++ {
		fmt.Println(fibonacci(i))
	}
	fmt.Println(fibonacci(posicao))
}