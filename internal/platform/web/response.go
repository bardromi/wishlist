package web

import (
	"encoding/json"
	"log"
	"net/http"
)

// RespondError sends an error reponse back to the client.
func RespondError(w http.ResponseWriter, err string, statusCode int) {
	log.Println(err)
	Respond(w, map[string]string{"message": err}, statusCode)
}

// Respond converts a Go value to JSON and sends it to the client.
// If code is StatusNoContent, v is expected to be nil.
func Respond(w http.ResponseWriter, data interface{}, statusCode int) {
	// If there is nothing to marshal then set status code and return.
	if statusCode == http.StatusNoContent {
		w.WriteHeader(statusCode)
	}

	// Convert the response value to JSON.
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// Set the content type and headers once we know marshaling has succeeded.
	w.Header().Set("Content-Type", "application/json")

	// Write the status code to the response.
	w.WriteHeader(statusCode)

	// Send the result back to the client.
	if _, err := w.Write(jsonData); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
