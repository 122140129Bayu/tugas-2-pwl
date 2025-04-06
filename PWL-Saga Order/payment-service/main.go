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

func makePaymentHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    response := Response{Success: false, Message: "Payment failed intentionally"}
    json.NewEncoder(w).Encode(response)
}

func refundPaymentHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    response := Response{Success: true, Message: "Payment refunded successfully"}
    json.NewEncoder(w).Encode(response)
}

func main() {
    http.HandleFunc("/make-payment", makePaymentHandler)
    http.HandleFunc("/refund-payment", refundPaymentHandler)

    log.Println("Payment Service running on port 8082...")
    log.Fatal(http.ListenAndServe(":8082", nil))
}
