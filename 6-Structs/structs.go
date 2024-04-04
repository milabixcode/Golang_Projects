package main

import "fmt"

//type: cria um tipo novo: tipo usuario
type usuario struct {
	nome string
	idade uint8
	endereco endereco
}

//Struct dentro de Struct
type endereco struct {
	logradouro string
	numero uint8
}

func main() {
	fmt.Println("Arquivo Structs")

	var u usuario 
	u.nome = "Davi"
	u.idade = 21
	fmt.Println(u)

	//usando inferencia de tipos

	enderecoExemplo := endereco{"Rua dos Bobos", 0}
	usuario2 := usuario{"Bia", 22, enderecoExemplo}
	fmt.Println(usuario2)

	//outra forma quando nao tenho todos os dados:
	usuario3 := usuario{nome: "Matheus"}
	fmt.Println(usuario3)

	usuario4 := usuario{idade: 30 }
	fmt.Println(usuario4)

	
}