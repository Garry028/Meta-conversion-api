package main

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// Define a struct to represent the destination settings
type DestinationSettings struct {
	PixelID       string `json:"pixelId"`
	AccessKey     string `json:"accessKey"`
	TestEventCode string `json:"testEventCode"`
}

// Define a struct to represent the events
type Event struct {
	EventName string `json:"eventName"`
}

// Define a struct to represent the entire payload
type Payload struct {
	DestinationSettings DestinationSettings `json:"destinationSettings"`
	Events              []Event             `json:"events"`
}

// Function to handle the HTTP POST request for Purchase Events
func handlePurchaseEvent(w http.ResponseWriter, r *http.Request) {
	// means only the post methid is allowed else throw error
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Decode the JSON payload into the Payload struct
	var payload Payload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	// Validate the required fields in destination settings
	if payload.DestinationSettings.PixelID == "" || payload.DestinationSettings.AccessKey == "" || payload.DestinationSettings.TestEventCode == "" {
		http.Error(w, "Incomplete destination settings data", http.StatusBadRequest)
		return
	}

	// Loop through the events and perform any required normalization and hashing
	for _, event := range payload.Events {
		// Do normalization and hashing for each event
		normalizedEventName := normalizeEventName(event.EventName)
		hashedEventName := hashEventName(normalizedEventName)

		// Simulate sending the normalized and hashed event data to Facebook
		fmt.Printf("Sending event to Facebook: %s\n", hashedEventName)
	}

	// Prepare the data to be sent to Facebook
	data := url.Values{}
	data.Set("pixelId", payload.DestinationSettings.PixelID)
	data.Set("accessKey", payload.DestinationSettings.AccessKey)
	data.Set("testEventCode", payload.DestinationSettings.TestEventCode)

	// Simulate sending the data to the test pixel on Facebook (replace with actual Facebook API call)
	// For this example, we will just print the data to the console.
	fmt.Println("Sending data to Facebook: ", data.Encode())

	// Provide a response back to the client
	response := map[string]interface{}{
		"status":  "success",
		"message": "Purchase events sent to Facebook",
	}
	jsonResponse, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

// Helper function to perform normalization on event names
func normalizeEventName(eventName string) string {
	//will convert the event name to lowercase.
	return strings.ToLower(eventName)
}

// Helper function to perform hashing on event names
func hashEventName(eventName string) string {
	//will use a simple hash function (SHA-1) from the crypto package.
	hasher := sha1.New()
	hasher.Write([]byte(eventName))
	hashed := hex.EncodeToString(hasher.Sum(nil))
	return hashed
}

func main() {
	// Set up a route for handling Purchase Events
	http.HandleFunc("/send/facebook", handlePurchaseEvent)

	// Start the web server
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
