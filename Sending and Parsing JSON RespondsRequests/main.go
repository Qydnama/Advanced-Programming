package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type RequestData struct {
	Message string `json:"message"`
}

type ResponseData struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var reqData RequestData
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&reqData)
	if err != nil {
		http.Error(w, "Error parsing JSON data", http.StatusBadRequest)
		return
	}

	if reqData.Message == "" {
		respondJSON(w, http.StatusBadRequest, ResponseData{
			Status:  "400",
			Message: "Invalid JSON message",
		})
		return
	}

	log.Printf("Message received: %s\n", reqData.Message)

	respondJSON(w, http.StatusOK, ResponseData{
		Status:  "success",
		Message: "Data successfully received",
	})
}

func respondJSON(w http.ResponseWriter, statusCode int, data ResponseData) {
	response, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(response)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	log.Print("starting server on port :8080")

	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
