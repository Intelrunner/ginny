package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// ResponseData represents the structure of the JSON response
type ResponseData struct {
	Message string            `json:"message"`
	Status  string            `json:"status"`
	Data    map[string]string `json:"data"`
}

// Handle the root URL
func homeHandler(w http.ResponseWriter, r *http.Request) {
	html := `<h1>🚀 Welcome to Your Go Server!</h1>
	<p>Go makes web development efficient and fast!</p>`
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, html)
}

// Handle the /api/data endpoint
func getDataHandler(w http.ResponseWriter, r *http.Request) {
	response := ResponseData{
		Message: "Hello, Go!",
		Status:  "success",
		Data: map[string]string{
			"framework": "Go",
			"version":   "1.x",
		},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Handle the /api/submit POST endpoint
func submitDataHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var requestData map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	response := map[string]interface{}{
		"message": "Data received!",
		"data":    requestData,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func main() {
	// Define the routes
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/api/data", getDataHandler)
	http.HandleFunc("/api/submit", submitDataHandler)

	// Start the server on port 8080
	fmt.Println("Server is running on http://0.0.0.0:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
