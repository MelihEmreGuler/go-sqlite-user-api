package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/MelihEmreGuler/go-sqlite-user-api/pkg/database"
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

	dataGet()

	// Create a new router using the gorilla/mux package
	r := mux.NewRouter()

	// Define routes for various endpoints
	r.HandleFunc("/", index)                         // root endpoint
	r.HandleFunc("/post", post)                      // post endpoint
	r.HandleFunc("/post/{category}/{id}", post)      // dynamic post endpoint with parameters
	r.HandleFunc("/product", product).Methods("GET") // product endpoint that only accepts GET requests
	//r.HandleFunc("/product", addProduct).Methods("POST") // product endpoint that only accepts POST requests

	r.HandleFunc("/api_get_user", api_get_user)
	r.HandleFunc("/api_add_user", api_add_user)
	r.HandleFunc("/api_update_user", api_update_user)
	r.HandleFunc("/api_delete_user", api_delete_user)

	// Start the HTTP server on port 8080 using the router
	http.ListenAndServe(":8080", r)
}

// Handler function for the root endpoint
func index(w http.ResponseWriter, r *http.Request) {
	// Write a response message containing information about the API and its usage
	fmt.Fprintln(w, "Welcome to the User API")
	fmt.Fprintln(w, "This API allows you to perform CRUD (Create, Read, Update, Delete) operations on users")

	// Provide examples of how to use the API
	fmt.Fprintln(w, "To get a list of all users, make a GET request to /api_get_user")
	fmt.Fprintln(w, "To add a new user, make a POST request to /api_add_user with userName and userPassword parameters")
	fmt.Fprintln(w, "To update an existing user, make a POST request to /api_update_user with userID, userName, and userPassword parameters")
	fmt.Fprintln(w, "To delete a user, make a POST request to /api_delete_user with userID parameter")
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
	w.Write(j) //fmt.Fprintf(w, string(j))
}

func api_get_user(w http.ResponseWriter, r *http.Request) {
	// Print the matching URL to the console
	fmt.Println(r.URL.Path)

	// Set the content type of the response to JSON
	w.Header().Set("content-Type", "application/json")

	// Marshal the UserList slice from the database package into JSON format
	j, err := json.Marshal(database.UserList)
	if err != nil {
		panic(err)
	}

	// Write the JSON-encoded response to the client
	fmt.Fprintln(w, string(j))
}

func api_add_user(w http.ResponseWriter, r *http.Request) {
	// Print the matching URL to the console
	fmt.Println(r.URL.Path)

	// Call the AddUser function from the database package with the values from the request parameters
	database.AddUser(r.FormValue("userName"), r.FormValue("userPassword"))

	// Write a response message containing the added user information
	fmt.Fprintln(w, "userName:", r.FormValue("userName"), "userPassword:", r.FormValue("userPassword"))
}

func api_update_user(w http.ResponseWriter, r *http.Request) {
	// Print the matching URL to the console
	fmt.Println(r.URL.Path)

	// Parse the userID parameter from the request into an integer
	id, err := strconv.ParseInt(r.FormValue("userID"), 10, 32)
	if err != nil {
		panic(err)
	}

	// Call the UpdateUser function from the database package with the values from the request parameters
	database.UpdateUser(int(id), r.FormValue("userName"), r.FormValue("userPassword"))

	// Write a response message containing the updated user information
	fmt.Fprintln(w, "UserID:", id, "userName:", r.FormValue("userName"), "userPassword:", r.FormValue("userPassword"))
}

func api_delete_user(w http.ResponseWriter, r *http.Request) {
	// Print the matching URL to the console
	fmt.Println(r.URL.Path)

	// Parse the userID parameter from the request into an integer
	id, err := strconv.ParseInt(r.FormValue("userID"), 10, 32)
	if err != nil {
		panic(err)
	}

	// Call the DeleteUser function from the database package with the parsed userID parameter
	database.DeleteUser(int(id))

	// Write a response message containing the deleted user information
	fmt.Fprintln(w, "UserID:", id)
}
