package main

import "fmt"

func main() {
	// Array é uma coleção de dados do mesmo tipo
	// Arrays possuem tamanho fixo, definido em tempo de compilação
	var filmes [5]string

	filmes[0] = "O Senhor dos Anéis"
	filmes[1] = "O Poderoso Chefão"
	filmes[2] = "Barbie"

	// Declaração curta
	numeros := [4]int{0, 2, 4, 6}
	fmt.Println(numeros)

	// Slices -> estruturas flexíveis, de tamanho dinâmico
	var novosNumeros []int

	novosNumeros = numeros[1:] // vai até o final do array
	novosNumeros = numeros[1:3] // vai até o elemento de índice 2
	fmt.Println(novosNumeros)

	numeros[2] = 7
	fmt.Println(novosNumeros)

	nomes := []string{"Ana", "Pedro"}
	fmt.Println(nomes)

	fmt.Println("----------------")
	for i := 0; i < len(numeros); i++ {
		fmt.Println(numeros[i])
	}

	// range -> enumerate em Python
	for indice, num := range numeros {
		fmt.Println("Índice:", indice, "- valor:", num)
	}

	fmt.Println("----------------")
	fmt.Println(numeros)
	modificarArray(numeros) // não altero o array original
	fmt.Println(numeros)

	fmt.Println(novosNumeros)
	modificarSlice(novosNumeros) // altera o slice original
	fmt.Println(novosNumeros)
	fmt.Println(numeros)
}

func modificarArray(a [4]int) {
	a[0] = 999
}

func modificarSlice(s []int) {
	s[0] = 999
}