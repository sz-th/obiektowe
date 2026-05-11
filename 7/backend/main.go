package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/mail"
	"strings"
	"time"
)

const (
	contentTypeJSON     = "application/json"
	headerContentType   = "Content-Type"
	methodNotAllowedMsg = "Method not allowed"
	invalidBodyMsg      = "Invalid request body"
	allowedOrigin       = "http://localhost:5173"
	serverAddr          = ":8080"

	maxBodyBytes      = 1 << 20
	maxFullNameLength = 200
	maxEmailLength    = 320
	maxAmount         = 1_000_000.0
	maxCartItems      = 1000
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

func setSecurityHeaders(w http.ResponseWriter) {
	h := w.Header()
	h.Set("X-Content-Type-Options", "nosniff")
	h.Set("X-Frame-Options", "DENY")
	h.Set("Referrer-Policy", "no-referrer")
}

func withCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		setSecurityHeaders(w)
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
	r.Body = http.MaxBytesReader(w, r.Body, maxBodyBytes)
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(dest); err != nil {
		http.Error(w, invalidBodyMsg, http.StatusBadRequest)
		return false
	}
	return true
}

func validatePayment(p PaymentRequest) error {
	name := strings.TrimSpace(p.FullName)
	if name == "" || len(name) > maxFullNameLength {
		return errors.New("invalid full name")
	}

	email := strings.TrimSpace(p.Email)
	if email == "" || len(email) > maxEmailLength || !strings.Contains(email, "@") {
		return errors.New("invalid email")
	}
	if _, err := mail.ParseAddress(email); err != nil {
		return errors.New("invalid email")
	}

	if p.Amount <= 0 || p.Amount > maxAmount {
		return errors.New("invalid amount")
	}
	return nil
}

func validateCart(c CartRequest) error {
	if len(c.Items) == 0 || len(c.Items) > maxCartItems {
		return errors.New("invalid items")
	}
	return nil
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
	if err := validatePayment(request); err != nil {
		http.Error(w, invalidBodyMsg, http.StatusBadRequest)
		return
	}

	message := fmt.Sprintf("Platnosc przyjeta na kwote %.2f", request.Amount)
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
	if err := validateCart(request); err != nil {
		http.Error(w, invalidBodyMsg, http.StatusBadRequest)
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

	srv := &http.Server{
		Addr:              serverAddr,
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       60 * time.Second,
		MaxHeaderBytes:    1 << 20,
	}

	log.Printf("Backend listening on %s", serverAddr)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
