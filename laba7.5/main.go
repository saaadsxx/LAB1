package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("Request: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
		log.Printf("Completed in %s", time.Since(start))
	})
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func dataHandler(w http.ResponseWriter, r *http.Request) {
	// Same as previous example
	fmt.Fprintln(w, "Data received")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/data", dataHandler)

	loggedMux := loggingMiddleware(mux)
	fmt.Println("HTTP Server with middleware is running on port 8080...")
	http.ListenAndServe(":8080", loggedMux)
}
