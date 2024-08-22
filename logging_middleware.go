package main

import (
	"log"
	"net/http"
	"time"
)

func LoggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Log request details
		log.Printf("Received request: %s %s", r.Method, r.URL.Path)

		// Call the next handler
		next.ServeHTTP(w, r)

		// Log response details
		duration := time.Since(start)
		log.Printf("Handled request: %s %s in %v", r.Method, r.URL.Path, duration)
	}
}
