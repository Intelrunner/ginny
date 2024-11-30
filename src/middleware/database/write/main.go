package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

// CardData represents the structure of the card data received from the HTTP request.
type CardData struct {
	CARD_TITLE   string `json:"CARD_TITLE"`
	MANA_COST    string `json:"MANA_COST"`
	EFFECT_TEXT  string `json:"EFFECT_TEXT"`
	SERIES       string `json:"SERIES"`
	// Add other fields as needed
}

os.Getenv("PROJECT_ID")


func main() {
	// Replace with your project
	// Get $PROJECT_ID and $SA_BUCKET $SA_CRED_FILE from the environment


	// Create a Firestore client
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID, option.WithCredentialsFile(serviceAccountKeyPath))
	if err != nil {
		log.Fatalf("Failed to create Firestore client: %v", err)
	}
	defer client.Close()

	// Define the HTTP handler
	http.HandleFunc("/add-card", func(w http.ResponseWriter, r *http.Request) {
		// Decode the JSON request body
		var cardData CardData
		if err := json.NewDecoder(r.Body).Decode(&cardData); err != nil {
			http.Error(w, fmt.Sprintf("Error decoding JSON: %v", err), http.StatusBadRequest)
			return
		}

		// Check if the card already exists in the "CARD_DB" collection
		docRef := client.Collection("CARD_DB").Doc(cardData.SERIES)
		docSnap, err := docRef.Get(ctx)
		if err != nil {
			// Handle errors getting the document
			http.Error(w, fmt.Sprintf("Error getting document: %v", err), http.StatusInternalServerError)
			return
		}

		// If the document exists, return an error
		if docSnap.Exists() {
			http.Error(w, fmt.Sprintf("Card with series '%s' already exists", cardData.SERIES), http.StatusConflict)
			return
		}

		// Create a new document in the "CARD_DB" collection
		_, err = client.Collection("CARD_DB").Doc(cardData.SERIES).Set(ctx, map[string]interface{}{
			"CARD_TITLE":   cardData.CARD_TITLE,
			"MANA_COST":    cardData.MANA_COST,
			"EFFECT_TEXT":  cardData.EFFECT_TEXT,
			// Add other fields as needed
		})
		if err != nil {
			http.Error(w, fmt.Sprintf("Error creating document: %v", err), http.StatusInternalServerError)
			return
		}

		// Respond with success
		fmt.Fprintf(w, "Card added successfully!")
	})

	// Start the HTTP server
	log.Fatal(http.ListenAndServe(":8080", nil))
}
