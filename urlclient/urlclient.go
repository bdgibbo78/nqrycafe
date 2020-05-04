package main

import (
	"bytes"
	"time"
	"encoding/json"
	"log"
	"net/http"

	"../pos"
	"../model"
)


func main() {

	// connect to the REST API
	baseURL := "http://localhost:8080/api/v1"

	c1 := pos.MakeCoffee("arabica", "latte", "large", []model.ItemType{"sugar"})

    //assertEqual(t, cost, float32(5.875))

	postData, lerr := json.Marshal(&c1)
	if lerr != nil {
		log.Fatalln(lerr)
	}
	request, err := http.NewRequest("POST", baseURL + "/order/coffee", bytes.NewBuffer(postData))
	if err != nil {
		log.Fatalln(err)
	}

	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	// Set the header information
	request.Header.Set("Content-Type", "application/json")

	// Get the response JSON
	response, err := client.Do(request)
	if err != nil {
		log.Fatalln("Failed to post Order: %e.\n", err)
	}

	defer response.Body.Close()
	decoder := json.NewDecoder(response.Body)

	var sale pos.Sale
	err = decoder.Decode(&sale)
	if err != nil {
		log.Fatalln("Failed to decode sale: ", err)
	}

	log.Printf("Coffee cost: %f", sale.Cost)
}
