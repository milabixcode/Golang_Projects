package main

import (
	"log"
	"net/http"
)

func main() {
	// HTTP É UM PROTOCOLO DE COMUNICAÇÃO - BASE DA COMUNICAÇÃO WEB

	// CLIENTE - SERVIDOR
	// Cliente faz uma requisição para o servidor, que processa a requisição
	// e devolve uma resposta

	// Request - Response

	// Rotas: maneira de identificar que tipo de mensagem está sendo enviada
	// e que tipo de processamento o servidor precisa fazer em cima dessa msg
		// URI - Identificador do recurso
		// Metodo - o que vc quer fazer com o recurso - GET, POST, PUT, DELETE

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Olá Mundo!"))
	})

	log.Fatal(http.ListenAndServe(":5000", nil))
}