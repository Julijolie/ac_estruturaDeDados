package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	var contatos [5]Contato

	for i := 0; i < len(contatos); i++ {
		fmt.Println("Nome:")
		nome, _ := leitor.ReadString('\n')

		fmt.Println("E-mail:")
		email, _ := leitor.ReadString('\n')

		contato := Contato{
			Nome:  nome,
			Email: email,
		}

		contatos = addContato(contato, contatos)
		fmt.Println("***************************************")
	}

	fmt.Println(contatos)
}

var leitor = bufio.NewReader(os.Stdin)

type Contato struct {
	Nome  string
	Email string
}

func addContato(contato Contato, contatos [5]Contato) [5]Contato {
	for i := range contatos {
		if contatos[i] == (Contato{}) {
			contatos[i] = contato
			break
		}
	}
	return contatos
}
