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

}

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