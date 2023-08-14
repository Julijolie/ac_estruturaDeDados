package main

import "fmt"

func main() {
	fmt.Println(soma(2, 5))
	fmt.Println(subtracao(4.0, 8.5))

	var x, y = 10.2, -2.45
	x, y = trocaValores(x, y)
	fmt.Println(x, y)

	usoAnonima()
}

// anônima (closure)
func usoAnonima() {
	dobra := func(x int) int {
		return 2 * x
	}

	resultado := dobra(5)
	fmt.Println(resultado)

	calcular := func(f func(int) int, x int) int {
		return f(x)
	}

	resultado = calcular(dobra, 3)
	fmt.Println(resultado)
}

// func nome(param1 tipo, param2 tipo, ...) tipoRetorno {}
func soma(a int, b int) int {
	return a + b
}

// parâmetros com mesmo tipo
func subtracao(a, b float64) float64 {
	return a - b
}

// múltiplos retornos
func trocaValores(a, b float64) (float64, float64) {
	return b, a
}