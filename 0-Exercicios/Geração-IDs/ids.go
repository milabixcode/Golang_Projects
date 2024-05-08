// Geração de IDs a partir de múltiplas go routines
// Implementar um service de gerar IDs incrementais(1,2,3,4)
// Implementar um controller com Múltiplas chamadas de go routine para gerar ID e imprimir o valor.
// Limitar as chamadas em 10 coroutines por vez
// Imprimir os IDs em ordem
// ** Emitir 1000 go routines no controller

package main

import (
	"fmt"
	"sync"
)

func main() {

	idInformado := 1

	canalDeIds := make(chan ID, 5000)

	var waitGroup sync.WaitGroup
	waitGroup.Add(1)

	// mutex é um tipo que permite acesso exclusivo a uma goroutine por vez
	//em um determinado trecho de código
	var mute sync.Mutex

	var idAtual ID = 2

	go func() {
		// 1 goRoutine que cria 5000 ids
		for i := 0; i < 5000; i++ {
			gerarIDs(ID(idInformado+i), canalDeIds)
		}
		waitGroup.Done()
	}()

	// Cria 10 goroutines concorrentes, cada uma recebe id do canal e imprime
	for i := 0; i < 10; i++ {
		waitGroup.Add(1)
		go func() {
			mensagem := <-canalDeIds

			mute.Lock()
			//se não for a ordem do id ser printado
			for idAtual != mensagem {
				// destrava para as outras goroutines poderem usar o idAtual 
				mute.Unlock()
				//travo novamente para checar se ainda é diferente
				mute.Lock()
			}

			fmt.Println(mensagem)
			idAtual += 1

			mute.Unlock()

			waitGroup.Done()
		}()
	}
	// Fala para a função esperar a contagem das goRoutines chegar em 0
	waitGroup.Wait()
}

//Struct de ID
type ID int

//Função que gera IDs incrementais e envia pelo canal
func gerarIDs(id ID, canal chan ID) {
	var nextID ID
	nextID = id + 1
	//envia o nextID para o canal
	canal <- nextID
}
