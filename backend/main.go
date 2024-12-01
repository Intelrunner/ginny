package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

// Product struct defines the product model
type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

//handler for the api/test endpoint
func test(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

// Handler for the /api/products endpoint
func getProducts(w http.ResponseWriter, r *http.Request) {
	// Sample product data
	products := []Product{
		{ID: 1, Name: "Product 1", Price: 10.0},
		{ID: 2, Name: "Product 2", Price: 20.0},
	}

	// Set response headers
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func main() {
	// Initialize the router
	r := mux.NewRouter()

	// Define routes
	r.HandleFunc("/api/products", getProducts).Methods("GET")
	r.HandleFunc("/api/test", test).Methods("GET")


	// Start the server
	const port = ":8080"
	log.Printf("Server running on port %s", port)
	log.Print(http.ListenAndServe(port, r))
}
