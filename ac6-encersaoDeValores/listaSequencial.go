package main

import (
	"fmt"
)

const M = 5

func main() {
	var lista [M]int
	var n = 0

	insereOrd(&lista, &n, 4)
	fmt.Println(lista)

	insereOrd(&lista, &n, 12)
	fmt.Println(lista)

	insereOrd(&lista, &n, 2)
	fmt.Println(lista)

	insereOrd(&lista, &n, 6)
	fmt.Println(lista)

	insereOrd(&lista, &n, 17)
	fmt.Println(lista)

	insereOrd(&lista, &n, 1)
	fmt.Println(lista)
}

func insereOrd(v *[M]int, n *int, novoValor int) {
	// Encontrar a posição correta para inserir o novoValor mantendo a lista ordenada.
	pos := 0
	for pos < *n && novoValor > v[pos] {
		pos++
	}

	// Deslocar os elementos à direita para abrir espaço para o novoValor.
	for i := *n; i > pos; i-- {
		v[i] = v[i-1]
	}

	// Inserir o novoValor na posição correta.
	v[pos] = novoValor

	// Incrementar o tamanho da lista.
	*n++
}
