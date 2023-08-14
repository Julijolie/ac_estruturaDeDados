package main

import "fmt"

type Pessoa struct {
	Nome string
	Idade int
	Altura float64
}
func main(){
	pessoa1 := Pessoa{
		Nome: "Juliana",
		Idade: 25,
		Altura: 1.64,
	}

	fmt.Println("Nome:", pessoa1.Nome)
    fmt.Println("Idade:", pessoa1.Idade)
    fmt.Println("Altura:", pessoa1.Altura)

}