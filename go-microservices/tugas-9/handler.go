package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	h8HelperRand "github.com/novalagung/gubrak/v2"
)

func PostRequest() {
	// initialize request data
	data := Data{Water: h8HelperRand.RandomInt(1, 100), Wind: h8HelperRand.RandomInt(1, 100)}
	apiURL := "https://jsonplaceholder.typicode.com/posts"

	reqJson, err := json.Marshal(data) // convert request data object to json
	client := &http.Client{}
	if err != nil {
		log.Fatalln(err.Error())
	}

	// make post request ot api url
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(reqJson))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		log.Fatalln(err.Error())
	}

	// check water and wind status based on request data
	statusWater := CheckWaterLevel(data.Water)
	statusWind := CheckWindLevel(data.Wind)
	res, err := client.Do(req) //send request to api
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer res.Body.Close()

	// read response body
	result := Data{}
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		log.Fatalln(err.Error())
	}

	resJson, _ := json.MarshalIndent(result, "", " ")

	// print response result
	fmt.Println(string(resJson))
	fmt.Println("status water:", statusWater)
	fmt.Println("status wind:", statusWind)
}