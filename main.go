package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	handlers := map[string]http.HandlerFunc{
		"GET /status":    LoggingMiddleware(statusHandler),
		"POST /add":      LoggingMiddleware(addHandler),
		"POST /subtract": LoggingMiddleware(subtractHandler),
		"POST /multiply": LoggingMiddleware(multiplyHandler),
		"POST /divide":   LoggingMiddleware(divideHandler),
	}

	port := ":8080"

	fmt.Printf("Listening on http://localhost%s...\n", port)
	for route, handler := range handlers {
		http.HandleFunc(route, handler)
		fmt.Println(route)
	}

	log.Fatal(http.ListenAndServe(port, nil))
}
