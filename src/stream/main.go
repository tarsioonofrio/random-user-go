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

type Request struct {
	Result []struct {
		Gender	string `json:"gender"`
		Email	string `json:"email"`
		Name  	struct {
			First string `json:"first"`
			Last  string `json:"last"`
		} `json:"name"`
		Login   struct {
			Username string `json:"username"`
			Password string `json:"password"`
		} `json:"login"`
		Location   struct {
			Name 	string `json:"street_name"`
			Number  string `json:"street_number"`
			City  	string `json:"city"`
			State  	string `json:"state"`
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

	var request Request
	err = json.Unmarshal(body, &request)
	if err != nil {
		panic(err.Error())
		//return r
	}

	fmt.Println(request)

	pessoa1 := &Pessoa{}
	endereco1 := &Endereco{}

	for _, r := range request.Result {
		//fmt.Printf("%s -> %s\n", i, r)
		pessoa1 = &Pessoa{
			Genero:		r.Gender,
			Email: 		r.Email,
			Nome:   	r.Name.First + " " + r.Name.Last,
			Username:	r.Login.Username,
			Password : 	r.Login.Password,

		}
		endereco1 = &Endereco{
			PessoaID:	pessoa1.ID,
			Rua:		r.Location.Name,
			Numero: 	r.Location.Number,
			Cidade: 	r.Location.City,
			Estado: 	r.Location.State,
		}
	}


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

	err = db.Insert(pessoa1)
	if err != nil {
		panic(err)
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