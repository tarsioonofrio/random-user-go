//package stream
// https://gist.github.com/josue/d5271bdfb36e1fad8e07b6ad9cd97629

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type User struct {
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

	var user User
	err = json.Unmarshal(body, &user)
	if err != nil {
		panic(err.Error())
		//return r
	}

	fmt.Println(user)

	//for key, result := range results {
	//	gender := result["gender"].(map[string]interface{})
	//	fmt.Println("Reading Value for Key :", key)
	//	//Reading each value by its key
	//	//fmt.Println("Id :", result["id"],
	//	//	"- Name :", result["name"],
	//	//	"- Department :", result["department"],
	//	//	"- Designation :", result["designation"])
	//	fmt.Println(gender)
	//	//fmt.Println("Address :", address["city"], address["state"], address["country"])
	//}
}