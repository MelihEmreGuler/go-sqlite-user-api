package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Define a Product struct to hold product information
type Product struct {
	ProductID   int    // use capital letter for field names to allow JSON encoding
	ProductName string // use capital letter for field names to allow JSON encoding
}

var apiProductList []Product // slice to store Product structs

// Define a function to populate apiProductList with sample data
func dataGet() {
	pro1 := Product{ProductID: 1, ProductName: "computer"}
	pro2 := Product{ProductID: 2, ProductName: "mobile phone"}

	apiProductList = append(apiProductList, pro1, pro2)
}

// Define the main API function to handle requests
func Api() {
	fmt.Println("api")

	// Create a new router using the gorilla/mux package
	r := mux.NewRouter()

	// Define routes for various endpoints
	r.HandleFunc("/", index)                           // root endpoint
	r.HandleFunc("/post", post)                        // post endpoint
	r.HandleFunc("/post/{category}/{id}", post)         // dynamic post endpoint with parameters
	r.HandleFunc("/product", product).Methods("GET")   // product endpoint that only accepts GET requests
	//r.HandleFunc("/product", addProduct).Methods("POST") // product endpoint that only accepts POST requests

	// Start the HTTP server on port 8080 using the router
	http.ListenAndServe(":8080", r)
}

// Handler function for the root endpoint
func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("index page"))
}

// Handler function for the post endpoints
func post(w http.ResponseWriter, r *http.Request) {
	// Extract parameters from the URL using the gorilla/mux package
	vars := mux.Vars(r)
	id := vars["id"]
	category := vars["category"]

	// Determine the request method (POST or GET)
	var a string
	if r.Method == "POST" {
		a = "post"
	} else if r.Method == "GET" {
		a = "get"
	}

	// Write a response message containing the extracted parameters and request method
	w.Write([]byte("post page, post id: " + id + " category: " + category + " method: " + a))
}

// Handler function for the product endpoint that returns a JSON-encoded list of products
func product(w http.ResponseWriter, r *http.Request) {
	// Print the matching URL to the console
	fmt.Println(r.URL.Path)

	// Marshal the apiProductList slice into JSON format
	j, err := json.Marshal(apiProductList)
	if err != nil {
		// Handle errors when marshaling to JSON format
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write the JSON-encoded response to the client
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}
