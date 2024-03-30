package main

import (
	"fmt"
)

type Pessoa struct {
	Nome       string
	Idade      int
	Cadastrado bool
}

func (p Pessoa) String() string {
	return fmt.Sprintf("%v", p.Nome)
}

func main() {
	var Pessoas = []Pessoa{{Nome: "Teste", Idade: 33, Cadastrado: true}, {Nome: "Sophia", Idade: 22, Cadastrado: false}}

	Pessoas = append(Pessoas, Pessoa{Nome: "Zé", Idade: 19, Cadastrado: false})
	Pessoas = append(Pessoas, Pessoa{Nome: "Joazinho", Idade: 8, Cadastrado: true})
	Pessoas = append(Pessoas, Pessoa{Nome: "Marcela", Idade: 69, Cadastrado: true})

	for _, Pessoa := range Pessoas {
		if Pessoa.Cadastrado {
			fmt.Println(Pessoa)
		}
	}
}
