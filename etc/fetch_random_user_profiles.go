// Fetches random user profiles and outputs their emails.
//
// @author - Josue Rodriguez <code@josuerodriguez.com>
//
// @date - Aug 23, 2019
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var apiURL = "https://randomuser.me/api/"
var defaultFetchTotalProfiles = 10

// JSON structure from API response
type person struct {
	Results []struct {
		Gender string `json:"gender"`
		Name   struct {
			Title string `json:"title"`
			First string `json:"first"`
			Last  string `json:"last"`
		} `json:"name"`
		Email string `json:"email"`
	} `json:"results"`
}

// result with Person & Error
type result struct {
	Person  person
	Error   error
	Latency time.Duration
}

func main() {
	startMain := time.Now()
	defer func() {
		fmt.Printf("\nTotal Spent: %v \n", time.Since(startMain))
	}()

	numbPtr := flag.Int("total", defaultFetchTotalProfiles, "total random profiles")
	flag.Parse()

	fmt.Printf("Fetching total profiles: %v\n\n", *numbPtr)

	// create channel to track create persons
	ch := make(chan result)

	// fetch up to X random persons
	go fetchUpToRandomPeople(*numbPtr, ch)

	// listen to channel for persons fetched
	for r := range ch {
		showPersonInfo(r)
	}
}

// output to STDOUT the result > person info
func showPersonInfo(r result) {
	if r.Error != nil {
		log.Printf("Response Error because err: %v\n", r.Error)
		return
	}

	if len(r.Person.Results) == 0 {
		log.Printf("Response Error: Person info not found\n")
		return
	}

	info := r.Person.Results[0]
	fmt.Printf("[%v] Email: %v \n", r.Latency, info.Email)
}

// fetch up to X random people and send back to channel
func fetchUpToRandomPeople(total int, ch chan result) {
	for i := 0; i <= total; i++ {
		ch <- getRandomUser()
		//showPersonInfo(getRandomUser())
	}

	close(ch)
}

// fetch random person from API url, then parse JSON and encode as 'person' struct.
func getRandomUser() result {
	startTime := time.Now()

	// get a random user from the API endpoint
	req, err := http.Get(apiURL)
	if err != nil {
		log.Fatal(err)
	}
	defer req.Body.Close()

	// init result
	var r result

	r.Latency = time.Since(startTime)

	if req.StatusCode != http.StatusOK {
		// catch none http-200 status code error
		r.Error = fmt.Errorf("API Error: Unable to get API request because HTTP status %d\n", req.StatusCode)
		return r
	}

	// read HTTP JSON body
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		// catch non-response error
		r.Error = fmt.Errorf("API Error: Unable to get API request because err: %v with HTTP status %d\n", err, req.StatusCode)
		return r
	}

	// convert JSON to struct
	err = json.Unmarshal(body, &r.Person)
	if err != nil {
		// capture parsing error
		r.Error = fmt.Errorf("JSON parse err: %v\n", err)
		return r
	}

	return r
}
