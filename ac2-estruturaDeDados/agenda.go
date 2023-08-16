package main

import (
	"fmt"
	"bufio"
	"os"
)


func main() {
	var nome string
	var email string

	fmt.Println("Nome: ")
	fmt.Scanln(&nome)

	fmt.Println("E-mail: ")
	fmt.Scanln(&email)

	leitor := bufio.NewReader(os.Stdin)
	msg, _ := leitor.ReadString('\n')
	fmt.Println(msg)

	var contatos [5]Contato
	var c = Contato{
		Nome : nome,
		Email : email,
	}

	contatos = addContato(c , contatos)

	contatos = excluirContato(contatos)

	fmt.Println(contatos)
}

type Contato struct {
	Nome string
	Email string
}

func addContato(contato Contato, contatos [5]Contato) [5]Contato {
	for i, c := range contatos {
		if (c == Contato{}) {
			contatos[i] = contato
			break
		}
	}
	return contatos
}

func excluirContato(contatos [5]Contato) [5]Contato {
	for i, c := range contatos{
		if (c == Contato{}){
			contatos [i - 1] = c
			break
		}
	}
	return contatos
}