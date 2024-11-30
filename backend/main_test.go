package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/gorilla/mux"

)

// TestGetProducts tests the /api/products endpoint
func TestGetProducts(t *testing.T) {
	// Create a request to pass to the handler
	req, err := http.NewRequest("GET", "/api/products", nil)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	// Create a ResponseRecorder to record the response
	recorder := httptest.NewRecorder()

	// Create a router and register the handler
	r := mux.NewRouter()
	r.HandleFunc("/api/products", getProducts).Methods("GET")

	// Serve the HTTP request
	r.ServeHTTP(recorder, req)

	// Check the status code
	assert.Equal(t, http.StatusOK, recorder.Code, "Expected status code 200")

	// Decode the response body
	var products []Product
	err = json.NewDecoder(recorder.Body).Decode(&products)
	if err != nil {
		t.Fatalf("Could not decode response: %v", err)
	}

	// Validate the response content
	expectedProducts := []Product{
		{ID: 1, Name: "Product 1", Price: 10.0},
		{ID: 2, Name: "Product 2", Price: 20.0},
	}
	assert.Equal(t, expectedProducts, products, "Response body does not match")
}
