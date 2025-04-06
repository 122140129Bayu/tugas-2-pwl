package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func createOrderHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Simulasi order berhasil
	response := Response{
		Success: true,
		Message: "Order created successfully",
	}
	json.NewEncoder(w).Encode(response)
}

func cancelOrderHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Simulasi pembatalan order
	response := Response{
		Success: true,
		Message: "Order cancelled successfully",
	}
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/create-order", createOrderHandler)
	http.HandleFunc("/cancel-order", cancelOrderHandler)

	log.Println("Order Service running on port 8081...")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}
}
