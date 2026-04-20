package main

import (
	"encoding/json"
	"log"
	"net/http"
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
	Total float64   `json:"total"`
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
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next(w, r)
	}
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func paymentsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var request PaymentRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(PaymentResponse{Message: "Platnosc przyjeta"})
}

func cartHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var request CartRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(PaymentResponse{Message: "Koszyk przyjety"})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/products", withCORS(productsHandler))
	mux.HandleFunc("/api/cart", withCORS(cartHandler))
	mux.HandleFunc("/api/payments", withCORS(paymentsHandler))

	log.Println("Backend listening on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
