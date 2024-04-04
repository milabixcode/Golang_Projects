package main

import (
	"fmt"
	"math/rand"
	"sort"
)
func criaVetor(tamanho int) [] int {
	vetor := make([] int, tamanho)
	for i:= 0; i < tamanho; i++ {
		vetor[i] = rand.Intn(101)
	}
	return vetor
}

func bubbleSort(vetor[] int) [] int {
	n := len(vetor)
	for i:= 0; i < n-1; i++ {
		//a cada iteração o maior já fica na ultima posicao (-i)
		//a cada iteração preciso comparar o penultimo com o ultimo(-1)
		for j:= 0; j < n - i - 1; j++ {
			if vetor[j] > vetor[j+1] {
				vetor[j], vetor[j+1] = vetor[j+1], vetor[j]
			}
		}
	}
	return vetor
}

//OUTRAS MANEIRAS DE ORDENAR
func quickSort(vetor []int) {
    if len(vetor) < 2 {
        return
    }
    left, right := 0, len(vetor)-1
    pivot := (left + right)/2
    vetor[pivot], vetor[right] = vetor[right], vetor[pivot]
    for i := range vetor {
        if vetor[i] < vetor[right] {
            vetor[i], vetor[left] = vetor[left], vetor[i]
            left++
        }
    }
    vetor[left], vetor[right] = vetor[right], vetor[left]
    quickSort(vetor[:left])
    quickSort(vetor[left+1:])
}

//Merge Sort
// Implementação do MergeSort
func mergeSort(vetor []int) []int {
    if len(vetor) <= 1 {
        return vetor
    }

    mid := len(vetor) / 2
    left := mergeSort(vetor[:mid])
    right := mergeSort(vetor[mid:])

    return merge(left, right)
}

func merge(left, right []int) []int {
    size, i, j := len(left)+len(right), 0, 0
    resultado := make([]int, size)

    for k := 0; k < size; k++ {
        if i > len(left)-1 && j <= len(right)-1 {
            resultado[k] = right[j]
            j++
        } else if j > len(right)-1 && i <= len(left)-1 {
            resultado[k] = left[i]
            i++
        } else if left[i] < right[j] {
            resultado[k] = left[i]
            i++
        } else {
            resultado[k] = right[j]
            j++
        }
    }
    return resultado
}

//HEAP SORT
// Implementação do HeapSort
func heapSort(vetor []int) {
    n := len(vetor)

    // Constrói um heap máximo
    // Começando pelo último nó não-folha e indo até a raiz
    for i := n/2 - 1; i >= 0; i-- {
        heapify(vetor, n, i)
    }

    // Extrai elementos um por um do heap
    for i := n - 1; i > 0; i-- {
        // Move a raiz atual para o final
        vetor[0], vetor[i] = vetor[i], vetor[0]
        // Chama heapify na pilha reduzida
        heapify(vetor, i, 0)
    }
}

// Função para transformar uma subárvore em um heap máximo
func heapify(vetor []int, tamanho, raiz int) {
    maior := raiz // Inicializa o maior como raiz
    esquerda := 2*raiz + 1  // Índice do filho esquerdo
    direita := 2*raiz + 2   // Índice do filho direito

    // Se o filho esquerdo é maior que a raiz
    if esquerda < tamanho && vetor[esquerda] > vetor[maior] {
        maior = esquerda
    }

    // Se o filho direito é maior que o maior até agora
    if direita < tamanho && vetor[direita] > vetor[maior] {
        maior = direita
    }

    // Se o maior não é a raiz
    if maior != raiz {
        // Troca a raiz com o maior
        vetor[raiz], vetor[maior] = vetor[maior], vetor[raiz]
        // Recursivamente Heapify a subárvore afetada
        heapify(vetor, tamanho, maior)
    }
}


func main() {
	fmt.Println("Vetor Original")

	//Cria o vetor
	vetor := criaVetor(5)

	//Imprime o vetor
	fmt.Println(vetor)

	//Ordena vetor
	//BUBBLE SORT
	fmt.Println("Bubble Sort")
	vetorOrdenado := bubbleSort(vetor)

	//Imprime o vetor ordenado
	fmt.Println(vetorOrdenado)

	//FUNÇÃO SORT
	fmt.Println("Funçao Sort")
	sort.Ints(vetor)
	//Imprime o vetor ordenado
	fmt.Println(vetor)

	//QUICKSORT
	fmt.Println("Quick Sort")
	quickSort(vetor)
	fmt.Println(vetor)

	//MERGESORT
	fmt.Println("Merge Sort")
	mergeSort(vetor)
	fmt.Println(vetor)

	//HEAPSORT
	fmt.Println("Heap Sort")
	heapSort(vetor)
	fmt.Println(vetor)

}