//package stream
// https://gist.github.com/josue/d5271bdfb36e1fad8e07b6ad9cd97629

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/go-pg/pg"
	"github.com/joho/godotenv"

)


type UserJson struct {
	Results []struct {
		Genero string `json:"gender"`
		Email string `json:"email"`
		Name   struct {
			Primeiro string `json:"first"`
			Ultimo  string `json:"last"`
		} `json:"name"`
		Login   struct {
			Username string `json:"username"`
			Password  string `json:"password"`
		} `json:"login"`
		Location   struct {
			Rua string `json:"street_name"`
			Numero  string `json:"street_number"`
			Cidade  string `json:"city"`
			Estado  string `json:"state"`
		} `json:"location"`
	} `json:"results"`
}


func main() {

	// json data
	url := "https://randomuser.me/api/"

	res, err := http.Get(url)

	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err.Error())
	}

	var user UserJson
	err = json.Unmarshal(body, &user)
	if err != nil {
		panic(err.Error())
		//return r
	}

	fmt.Println(user)

	err = godotenv.Load("/home/tarsio/personal/randomuser-go/src/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	opt, err := pg.ParseURL(os.Getenv("PGCONN2"))
	if err != nil {
		panic(err)
	}

	db := pg.Connect(opt)
	defer db.Close()



	pessoa1 := &Pessoa{
		Nome:   "OIOI",
		Genero: "Masculino",
	}
	err = db.Insert(pessoa1)
	if err != nil {
		panic(err)
	}

	err = db.Insert(&Pessoa{
		Nome:   "OLAOLA",
		Genero: "Feminino",
	})
	if err != nil {
		panic(err)
	}

	endereco1 := &Endereco{
		Rua:    "Cool story",
		PessoaID: pessoa1.ID,
	}
	err = db.Insert(endereco1)
	if err != nil {
		panic(err)
	}

	// Select user by primary key.
	pessoa := &Pessoa{ID: pessoa1.ID}
	err = db.Select(pessoa)
	if err != nil {
		panic(err)
	}

	// Select all users.
	var pessoas []Pessoa
	err = db.Model(&pessoas).Select()
	if err != nil {
		panic(err)
	}

	fmt.Println(pessoa)
	fmt.Println(pessoas)

}