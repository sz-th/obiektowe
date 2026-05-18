package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

const (
	contentTypeJSON     = "application/json"
	headerContentType   = "Content-Type"
	methodNotAllowedMsg = "Method not allowed"
	invalidBodyMsg      = "Invalid request body"
	defaultOrigin       = "http://localhost:5173"
	defaultServerAddr   = "127.0.0.1:8080"

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

func envOrDefault(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

func setSecurityHeaders(w http.ResponseWriter) {
	h := w.Header()
	h.Set("X-Content-Type-Options", "nosniff")
	h.Set("X-Frame-Options", "DENY")
	h.Set("Referrer-Policy", "no-referrer")
	h.Set("Content-Security-Policy", "default-src 'none'; frame-ancestors 'none'")
	h.Set("Permissions-Policy", "geolocation=(), microphone=(), camera=()")
}

func withCORS(allowedOrigin string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		setSecurityHeaders(w)
		w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Vary", "Origin")
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

func main() {
	allowedOrigin := envOrDefault("ALLOWED_ORIGIN", defaultOrigin)
	serverAddr := envOrDefault("SERVER_ADDR", defaultServerAddr)

	mux := http.NewServeMux()
	mux.HandleFunc("/api/products", withCORS(allowedOrigin, productsHandler))
	mux.HandleFunc("/api/cart", withCORS(allowedOrigin, cartHandler))
	mux.HandleFunc("/api/payments", withCORS(allowedOrigin, paymentsHandler))

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
