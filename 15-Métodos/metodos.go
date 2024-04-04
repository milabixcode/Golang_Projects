package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvarUsuario() {
	fmt.Printf("Salvando os dados do usuário %s no banco de dados \n", u.nome)
}

func (u *usuario) fazerAniversario() {
	u.idade++
}
func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

func main() {
	usuario1 := usuario{"Davi", 17}
	usuario1.salvarUsuario()
	ehMaior := usuario1.maiorDeIdade()
	fmt.Println(ehMaior)

	usuario1.fazerAniversario()
	ehMaior = usuario1.maiorDeIdade()
	fmt.Println(ehMaior)
}
