//package database

package main

import (
	"fmt"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"log"
	"os"
	//"path"
	//"path/filepath"

	"github.com/joho/godotenv"
)


func main() {
	//gotenv.Load("../.env")
	err := godotenv.Load("/home/tarsio/personal/randomuser-go/src/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	opt, err := pg.ParseURL(os.Getenv("PGCONN2"))
	if err != nil {
		panic(err)
	}

	db := pg.Connect(opt)
	defer db.Close()
	err = createSchema(db)
	if err != nil {
		//fmt.Println(err)
		panic(err)
	}

	//pessoa1 := &Pessoa{
	//	Nome:   "OIOIOI",
	//	Genero: "Masculino",
	//}
	//err = db.Insert(pessoa1)
	//if err != nil {
	//	panic(err)
	//}
	//
	//err = db.Insert(&Pessoa{
	//	Nome:   "OLAOLAOLA",
	//	Genero: "Feminino",
	//})
	//if err != nil {
	//	panic(err)
	//}
	//
	//endereco1 := &Endereco{
	//	Rua:    "Cool story",
	//	PessoaID: pessoa1.ID,
	//}
	//err = db.Insert(endereco1)
	//if err != nil {
	//	panic(err)
	//}
	//
	//// Select user by primary key.
	//pessoa := &Pessoa{ID: pessoa1.ID}
	//err = db.Select(pessoa)
	//if err != nil {
	//	panic(err)
	//}
	//
	//// Select all users.
	//var pessoas []Pessoa
	//err = db.Model(&pessoas).Select()
	//if err != nil {
	//	panic(err)
	//}

	// Select story and associated author in one query.
	//endereco := new(Endereco)
	//err = db.Model(endereco).
	//	Relation("Nome").
	//	Where("endereco.id = ?", endereco1.ID).
	//	Select()
	//if err != nil {
	//	panic(err)
	//}

	fmt.Println(pessoa)
	fmt.Println(pessoas)
	//fmt.Println(story)
	// Output: User<1 admin [admin1@admin admin2@admin]>
	// [User<1 admin [admin1@admin admin2@admin]> User<2 root [root1@root root2@root]>]
	// Story<1 Cool story User<1 admin [admin1@admin admin2@admin]>>
	//for {
	//}
}

//func (s *Story) String() string {
//	return fmt.Sprintf("Story<%d %s %s>", s.ID, s.Title, s.Author)
//}

// createSchema creates database schema for User and Story models.
func createSchema(db *pg.DB) error {
	models := []interface{}{
		(*Pessoa)(nil),
		(*Endereco)(nil),
	}

	for _, model := range models {
		err := db.CreateTable(model, &orm.CreateTableOptions{
			//Temp: true, // temp table
			IfNotExists: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}