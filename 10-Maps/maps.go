package main

import "fmt"

func main() {
	fmt.Println("Maps")

	usuario := map[string]string {
		"nome" :		"Pedro",
		"sobrenome" :	 "Silva",
	}
	fmt.Println((usuario))

	//MAP QUE AS CHAVES SAO DO tIPO STRING E O VALOR É UM MAP E AS CHAVES SAO DO TIPO STRING E VALORES DO TIPO STRING
	usuario2 := map[string]map[string]string  {
		"nome" : {
			"primeiro": "João",
			"ultimo" : "Pedro",
		},
	}
	fmt.Println(usuario2)
	delete(usuario2, "nome")
	fmt.Println(usuario2)
}