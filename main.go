package main

import (
	"gateway_go/eureka"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	err := eureka.RegisterWithEureka()
	if err != nil {
		return
	}

	r := mux.NewRouter()

	// Define your routes
	r.HandleFunc("/service1", service1Handler).Methods("GET")
	r.HandleFunc("/service2", service2Handler).Methods("GET")

	// Start the server
	log.Println("Starting API Gateway on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func service1Handler(w http.ResponseWriter, r *http.Request) {
	// Proxy request to service1
	http.Redirect(w, r, "http://localhost:8081/service1", http.StatusTemporaryRedirect)
}

func service2Handler(w http.ResponseWriter, r *http.Request) {
	// Proxy request to service2
	http.Redirect(w, r, "http://localhost:8082/service2", http.StatusTemporaryRedirect)
}
