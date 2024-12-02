package main

import (
	"context"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/gorilla/mux"
)

// Environment variables for Firestore configuration
var (
	db         string
	projectID  string
)

// Product represents a sample product structure
type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func init() {
	// Load and validate environment variables
	db = os.Getenv("COLLECTION_ID")
	if db == "" {
		log.Fatalf("Environment variable COLLECTION_ID is not set")
	}

	projectID = os.Getenv("PROJECT_ID")
	if projectID == "" {
		log.Fatalf("Environment variable PROJECT_ID is not set")
	}

	// Seed random number generator
	rand.Seed(time.Now().UnixNano())
}

// addRecordHandler handles adding a document to Firestore.
func addRecordHandler(client *firestore.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()

		// Define the document to be added to the Firestore collection
		docData := map[string]interface{}{
			"name":    "Los Angeles",
			"state":   "CA",
			"country": "USA",
		}

		// Generate a unique document ID
		docID := generateRandomID()

		// Attempt to write the document to the Firestore collection
		_, err := client.Collection(db).Doc(docID).Set(ctx, docData)
		if err != nil {
			http.Error(w, "Failed to add document to Firestore", http.StatusInternalServerError)
			log.Printf("Firestore write error: %v", err)
			return
		}

		// Respond with success
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("Document added to Firestore successfully"))
		log.Printf("Document with ID %s added to Firestore successfully", docID)
	}
}

// generateRandomID generates a unique string ID for Firestore documents
func generateRandomID() string {
	return strconv.Itoa(rand.Intn(1_000_000_000)) // Random integer between 0 and 999,999,999
}

// getProducts handles the /api/products endpoint
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

// testHandler is a placeholder handler for the /api/test endpoint
func testHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Test endpoint working"))
}

func main() {
	// Set up Firestore client
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create Firestore client: %v", err)
	}
	defer client.Close()

	// Initialize the router
	r := mux.NewRouter()

	// Define routes
	r.HandleFunc("/api/products", getProducts).Methods("GET")
	r.HandleFunc("/api/test", testHandler).Methods("GET")
	r.Handle("/api/addRecord", addRecordHandler(client)).Methods("POST")

	// Start the server
	const port = ":8080"
	log.Printf("Server running on port %s", port)
	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
