package main

import "fmt"

// func funcao1() {
// 	fmt.Println("Executando a funcao 1")
// }

// func funcao2() {
// 	fmt.Println("Executando a funcao 2")
// }

func alunoEstaAprovado(n1, n2 float32) bool {
	defer fmt.Println("Média calculada. resultado será retornado.")	
	fmt.Println("Entrando na função para verificar se o aluno está aprovado")
	
	media := (n1 + n2)/2

	if media > 6 {
		return true
	}
	return false
}

func main() {
	//defer adia a execução da função
	//DEFER = ADIAR
	// defer funcao1()
	// funcao2()

	fmt.Println(alunoEstaAprovado(7,8))

}