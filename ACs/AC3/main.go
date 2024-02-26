package main

import (
	"fmt"
)

// Definindo o struct Contato
type Contato struct {
	Nome  string
	Email string
}

// Método para alterar o e-mail
func (c *Contato) AlterarEmail(nEmail string) {
	c.Email = nEmail
}

// Função para adicionar um Contato ao array
func adicionarContato(contato Contato, arrayContatos []Contato) []Contato {
	for i := range arrayContatos {
		if arrayContatos[i].Nome == "" && arrayContatos[i].Email == "" {
			arrayContatos[i] = contato
			break
		}
	}
	return arrayContatos
}
func excluirContato(arrayContatos []Contato) []Contato {
	for i := len(arrayContatos) - 1; i >= 0; i-- {
		if arrayContatos[i].Nome != "" || arrayContatos[i].Email != "" {
			arrayContatos[i] = Contato{} // Zerando o último contato válido
			return arrayContatos
		}
	}
	// Se não houver contato válido, retorna o array sem modificações
	return arrayContatos
}
func main() {
	// Criando um array de 5 elementos
	contatos := make([]Contato, 5)

	fmt.Println("\n1 - Adicionar Contato")
	fmt.Println("2 - Excluir Último Contato")
	fmt.Println("3 - Sair")
	fmt.Print("Escolha uma opção: ")

	var opcao int
	fmt.Scan(&opcao)
	for {

		switch opcao {
		case 1:
			fmt.Println("\n Adicione Contato")
			var novoContato Contato
			fmt.Print("Nome: ")
			fmt.Scan(&novoContato.Nome)
			fmt.Print("Email: ")
			fmt.Scan(&novoContato.Email)
			contatos = adicionarContato(novoContato, contatos)
			fmt.Println("Contato adicionado com sucesso!")
			fmt.Println("\n\nLista de Contatos:")
			for i, contato := range contatos {
				if contatos[i].Nome == "" {
					break
				} else {
					fmt.Printf("%d - Nome: %s, Email: %s\n", i+1, contato.Nome, contato.Email)
				}
			}

		case 2:
			contatos = excluirContato(contatos)
			fmt.Println("exclusão feita com sucesso")

			fmt.Println("\nLista de Contatos:")
			for i, contato := range contatos {
				if contatos[i].Nome == "" {
					break
				} else {
					fmt.Printf("%d - Nome: %s, Email: %s\n", i+1, contato.Nome, contato.Email)
				}
			}
		case 3:
			break
		}
	}

}
