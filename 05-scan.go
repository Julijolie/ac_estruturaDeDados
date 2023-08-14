package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	var x, z int
	var y float64

	fmt.Println("Informe dois inteiros")
	fmt.Scanln(&x, &z)
	fmt.Println(x)
	fmt.Println(z)

	fmt.Println("Informe um float")
	fmt.Scanln(&y)
	fmt.Println(y)

	leitor := bufio.NewReader(os.Stdin)
	fmt.Println("Escreva uma mensagem:")
	msg, _ := leitor.ReadString('\n')

	fmt.Println(msg)
}