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

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
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

	pessoa := &Pessoa{}
	endereco := &Endereco{}

	for _, r := range request.Result {
		//fmt.Printf("%s -> %s\n", i, r)
		pessoa = &Pessoa{
			Genero:		r.Gender,
			Email: 		r.Email,
			Nome:   	r.Name.First + " " + r.Name.Last,
			Username:	r.Login.Username,
			Password : 	r.Login.Password,

		}
		endereco = &Endereco{
			PessoaID: pessoa.ID,
			Rua:      r.Location.Name,
			Numero:   r.Location.Number,
			Cidade:   r.Location.City,
			Estado:   r.Location.State,
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

	err = db.Insert(pessoa)
	if err != nil {
		panic(err)
	}

	err = db.Insert(endereco)
	if err != nil {
		panic(err)
	}

	// Select user by primary key.
	pessoa = &Pessoa{ID: pessoa.ID}
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

	//fmt.Println(pessoa)
	//fmt.Println(pessoas)


	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials
	// and region from the shared configuration file ~/.aws/config.
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		Config: aws.Config{Region: aws.String("us-east-1")},
	}))

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	for _, r := range request.Result {
		//fmt.Printf("%s -> %s\n", i, r)
		av, err := dynamodbattribute.MarshalMap(r)
		//av["id"] = "1"
		if err != nil {
			fmt.Println("Got error marshalling new movie item:")
			fmt.Println(err.Error())
			os.Exit(1)
		}

		// Create a struct with the movie data and marshall that data into a map of AttributeValue objects.
		input := &dynamodb.PutItemInput{
			Item:      av,
			TableName: aws.String("randomuser"),
		}

		// Create the input for PutItem and call it. If an error occurs, print the error and exit. If no error occurs, print an message that the item was added to the table.
		_, err = svc.PutItem(input)
		if err != nil {
			fmt.Println("Got error calling PutItem:")
			fmt.Println(err.Error())
			os.Exit(1)
		}

		fmt.Println("Successfully added")

	}

}