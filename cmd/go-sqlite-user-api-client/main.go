package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "http://localhost:8080/product"

	// Product struct for unmarshalling JSON data from HTTP response
	type Product struct {
		ProductID   int // use capital letter for field names to allow JSON encoding
		ProductName string
	}

	// Make GET request to API endpoint
	resp, err := http.Get(url)
	errorCheck(err)

	// Read response body and store as byte slice
	jsonDataFromHttp, err := ioutil.ReadAll(resp.Body)
	errorCheck(err)

	// Create slice of Product structs to hold JSON data
	var jsonData []Product

	// Unmarshal JSON data into slice of Product structs
	err = json.Unmarshal([]byte(jsonDataFromHttp), &jsonData)
	errorCheck(err)

	// Print slice of Product structs to console
	fmt.Println(jsonData)
}

// Error checking function
func errorCheck(err error) {
	if err != nil {
		panic(err)
	}
}