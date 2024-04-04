package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e Slices")

	//Array: lista de valores de mesmo tipo
	//tamanho fixo e tipo de dados fixo

	var array1 [5]string
	array1[0] = "Posição 1"
	fmt.Println(array1)

	array2 := [5]string{"Posição 1", "Posição 2", "Posição 3", "Posição 4", "Posição 5"}
	fmt.Println(array2)

	// ... fixam o tamanho do array de acordo com a quantidade de itens passado
	array3 := [...] int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	//Bastante usado
	//Slice: tamanho dinâmico, tb tem limitação de tipo
	slice := []int{ 10, 11, 12, 13, 14, 15, 16, 17}
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array3))

	slice = append(slice, 18)
	fmt.Println(slice)

	//pedaço de array - primeiro numero no colchete :inclusivo e 
	//ultimo numero no colchete: exclusivo
	slice2 := array2[1:3]
	fmt.Println(slice2)

	array2[1] = "Posição Alterada"
	fmt.Println(slice2)

	//Arrays internos

	//função make: Aloca espaço na memória
	//tipo, tamanho, capacidade
	fmt.Println("-------------")
	//array interno: cria um array de 15 posuções e retorna um slice de 10 posições
	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) //tamanho
	fmt.Println(cap(slice3)) //capacidade

	slice3 = append(slice3, 5)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) //tamanho
	fmt.Println(cap(slice3)) //capacidade

	//O Go quando vê que a capacidade do slice vai estourar o tamanho ele cria um outro array
	//e dobra o tamanho dele
	//Por isso o slice nunca tem um tamanho limite
	slice3 = append(slice3, 6)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) //tamanho
	fmt.Println(cap(slice3)) //capacidade

	//Quando não passa o último parâmetro no make ele entende que a capacidade é igual ao tamanho
	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	fmt.Println(len(slice4)) //tamanho
	fmt.Println(cap(slice4)) //capacidade

	//Dobra o tamanho
	slice4 = append(slice4, 10)
	fmt.Println(slice4)
	fmt.Println(len(slice4)) //tamanho
	fmt.Println(cap(slice4)) //capacidade

}
