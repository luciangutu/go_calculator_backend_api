package main

import (
	"encoding/json"
	"net/http"
)

type StatusResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type ResponseBody struct {
	Result int `json:"result"`
}

type RequestBody map[string]int

func statusHandler(w http.ResponseWriter, r *http.Request) {
	response := StatusResponse{
		Status:  "OK",
		Message: "Service is running",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	var requestBody RequestBody
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	sum := 0
	for _, value := range requestBody {
		sum += value
	}

	response := ResponseBody{Result: sum}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func subtractHandler(w http.ResponseWriter, r *http.Request) {
	var requestBody RequestBody
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if len(requestBody) == 0 {
		http.Error(w, "No numbers provided", http.StatusBadRequest)
		return
	}

	sub := 0

	for key, value := range requestBody {
		sub = value
		delete(requestBody, key)
		break
	}

	for _, value := range requestBody {
		sub -= value
	}

	response := ResponseBody{Result: sub}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func multiplyHandler(w http.ResponseWriter, r *http.Request) {
	var requestBody RequestBody
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if len(requestBody) == 0 {
		http.Error(w, "No numbers provided", http.StatusBadRequest)
		return
	}

	result := 1
	for _, value := range requestBody {
		result = result * value
	}

	response := ResponseBody{Result: result}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func divideHandler(w http.ResponseWriter, r *http.Request) {
	var requestBody RequestBody
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if len(requestBody) == 0 {
		http.Error(w, "No numbers provided", http.StatusBadRequest)
		return
	}

	for _, value := range requestBody {
		if value == 0 {
			http.Error(w, "Division by zero is not allowed", http.StatusBadRequest)
			return
		}
	}

	result := 0

	for key, value := range requestBody {
		result = value
		delete(requestBody, key)
		break
	}

	for _, value := range requestBody {
		result = result / value
	}

	response := ResponseBody{Result: result}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
