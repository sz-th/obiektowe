package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const (
	contentTypeJSON     = "application/json"
	headerContentType   = "Content-Type"
	methodNotAllowedMsg = "Method not allowed"
	invalidBodyMsg      = "Invalid request body"
	allowedOrigin       = "http://localhost:5173"
	serverAddr          = ":8080"
)

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type PaymentRequest struct {
	FullName string  `json:"fullName"`
	Email    string  `json:"email"`
	Amount   float64 `json:"amount"`
}

type CartRequest struct {
	Items []Product `json:"items"`
}

type PaymentResponse struct {
	Message string `json:"message"`
}

var products = []Product{
	{ID: 1, Name: "Laptop", Price: 3999.99},
	{ID: 2, Name: "Mysz", Price: 149.99},
	{ID: 3, Name: "Klawiatura", Price: 249.99},
}

func withCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next(w, r)
	}
}

func requireMethod(w http.ResponseWriter, r *http.Request, method string) bool {
	if r.Method != method {
		http.Error(w, methodNotAllowedMsg, http.StatusMethodNotAllowed)
		return false
	}
	return true
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set(headerContentType, contentTypeJSON)
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		log.Printf("encode response: %v", err)
	}
}

func decodeJSON(w http.ResponseWriter, r *http.Request, dest any) bool {
	if err := json.NewDecoder(r.Body).Decode(dest); err != nil {
		http.Error(w, invalidBodyMsg, http.StatusBadRequest)
		return false
	}
	return true
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	if !requireMethod(w, r, http.MethodGet) {
		return
	}
	writeJSON(w, http.StatusOK, products)
}

func paymentsHandler(w http.ResponseWriter, r *http.Request) {
	if !requireMethod(w, r, http.MethodPost) {
		return
	}
	defer r.Body.Close()

	var request PaymentRequest
	if !decodeJSON(w, r, &request) {
		return
	}

	message := fmt.Sprintf("Platnosc przyjeta dla %s na kwote %.2f", request.FullName, request.Amount)
	writeJSON(w, http.StatusOK, PaymentResponse{Message: message})
}

func cartHandler(w http.ResponseWriter, r *http.Request) {
	if !requireMethod(w, r, http.MethodPost) {
		return
	}
	defer r.Body.Close()

	var request CartRequest
	if !decodeJSON(w, r, &request) {
		return
	}

	message := fmt.Sprintf("Koszyk przyjety (%d pozycji)", len(request.Items))
	writeJSON(w, http.StatusOK, PaymentResponse{Message: message})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/products", withCORS(productsHandler))
	mux.HandleFunc("/api/cart", withCORS(cartHandler))
	mux.HandleFunc("/api/payments", withCORS(paymentsHandler))

	log.Printf("Backend listening on %s", serverAddr)
	if err := http.ListenAndServe(serverAddr, mux); err != nil {
		log.Fatal(err)
	}
}
