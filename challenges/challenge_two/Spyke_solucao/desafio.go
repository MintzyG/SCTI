package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

type Pessoa struct {
	sync.Mutex
	Id         int
	Nome       string
	Idade      int
	Cadastrado bool
}

var Pessoas = []Pessoa{{Id: 10, Nome: "Teste", Idade: 33, Cadastrado: true}, {Id: 11, Nome: "Sophia", Idade: 22, Cadastrado: false}}
var N_Id = 11

func main() {
	fmt.Printf("Starting server at port 8080\n")

	var fileServer = http.FileServer((http.Dir("./static")))
	http.Handle("/", fileServer)

	http.HandleFunc("/app/user/create-user", Create)
	http.HandleFunc("/app/user/list-users", List)
	http.HandleFunc("/app/user/registered-users", Registered)

	http.ListenAndServe(":8080", nil)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func Create(w http.ResponseWriter, r *http.Request) {
	var Nova_pessoa Pessoa

	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method not allowed: %s\n", r.Method)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&Nova_pessoa)
	if err != nil {
		log.Fatal(err)
		return
	}

	N_Id++
	Nova_pessoa.Id = N_Id
	Pessoas = append(Pessoas, Nova_pessoa)
}

func List(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	Resp, err := json.Marshal(Pessoas)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(Resp)
}

func Registered(w http.ResponseWriter, r *http.Request) {
	var C_pessoas []Pessoa
	for _, Pessoa := range Pessoas {
		if Pessoa.Cadastrado {
			C_pessoas = append(C_pessoas, Pessoa)
		}
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	Resp, err := json.Marshal(C_pessoas)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(Resp)
}
