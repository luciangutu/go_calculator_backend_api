package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAddHandler(t *testing.T) {
	requestBody := `{"number1": 5, "number2": 10}`
	req := httptest.NewRequest("POST", "/add", bytes.NewBufferString(requestBody))
	w := httptest.NewRecorder()

	addHandler(w, req)

	res := w.Result()
	if res.StatusCode != http.StatusOK {
		t.Fatalf("Expected status code 200, got %v", res.StatusCode)
	}

	var response ResponseBody
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if response.Result != 15 {
		t.Errorf("Expected result to be 15, got %d", response.Result)
	}
}

func TestSubtractHandler(t *testing.T) {
	requestBody := map[string]int{"a": 10, "b": 3, "c": 2}
	body, _ := json.Marshal(requestBody)
	req := httptest.NewRequest(http.MethodPost, "/subtract", bytes.NewReader(body))
	recorder := httptest.NewRecorder()

	subtractHandler(recorder, req)

	resp := recorder.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, resp.StatusCode)
	}

	var responseBody map[string]int
	json.NewDecoder(resp.Body).Decode(&responseBody)
	if responseBody["result"] != 5 {
		t.Errorf("Expected result to be 5, got %d", responseBody["result"])
	}
}

func TestDivideHandler(t *testing.T) {
	requestBody := map[string]int{"a": 20, "b": 5, "c": 2}
	body, _ := json.Marshal(requestBody)
	req := httptest.NewRequest(http.MethodPost, "/divide", bytes.NewReader(body))
	recorder := httptest.NewRecorder()

	divideHandler(recorder, req)

	resp := recorder.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, resp.StatusCode)
	}

	var responseBody map[string]int
	json.NewDecoder(resp.Body).Decode(&responseBody)
	if responseBody["result"] != 2 {
		t.Errorf("Expected result to be 2, got %d", responseBody["result"])
	}
}

func TestMultiplyHandler(t *testing.T) {
	requestBody := `{"number1": 2, "number2": 3}`
	req := httptest.NewRequest("POST", "/multiply", bytes.NewBufferString(requestBody))
	w := httptest.NewRecorder()

	multiplyHandler(w, req)

	res := w.Result()
	if res.StatusCode != http.StatusOK {
		t.Fatalf("Expected status code 200, got %v", res.StatusCode)
	}

	var response ResponseBody
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if response.Result != 6 {
		t.Errorf("Expected result to be 6, got %d", response.Result)
	}
}
