package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/mail"
	"strings"
)

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

func handlePostJSON[T any](
	w http.ResponseWriter,
	r *http.Request,
	validate func(T) error,
	buildMessage func(T) string,
) {
	if !requireMethod(w, r, http.MethodPost) {
		return
	}
	defer r.Body.Close()

	var request T
	if !decodeJSON(w, r, &request) {
		return
	}
	if err := validate(request); err != nil {
		http.Error(w, invalidBodyMsg, http.StatusBadRequest)
		return
	}

	writeJSON(w, http.StatusOK, PaymentResponse{Message: buildMessage(request)})
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

func validateProduct(item Product) error {
	name := strings.TrimSpace(item.Name)
	if item.ID <= 0 || name == "" || len(name) > maxFullNameLength {
		return errors.New("invalid product")
	}
	if item.Price <= 0 || item.Price > maxAmount {
		return errors.New("invalid product price")
	}
	return nil
}

func validateCart(c CartRequest) error {
	if len(c.Items) == 0 || len(c.Items) > maxCartItems {
		return errors.New("invalid items")
	}
	for _, item := range c.Items {
		if err := validateProduct(item); err != nil {
			return err
		}
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
	handlePostJSON(w, r, validatePayment, func(p PaymentRequest) string {
		return fmt.Sprintf("Platnosc przyjeta na kwote %.2f", p.Amount)
	})
}

func cartHandler(w http.ResponseWriter, r *http.Request) {
	handlePostJSON(w, r, validateCart, func(c CartRequest) string {
		return fmt.Sprintf("Koszyk przyjety (%d pozycji)", len(c.Items))
	})
}
