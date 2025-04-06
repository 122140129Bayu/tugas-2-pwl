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

func shipOrderHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    response := Response{Success: false, Message: "Shipping failed intentionally"}
    json.NewEncoder(w).Encode(response)
}

func cancelShippingHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    response := Response{Success: true, Message: "Shipping cancelled successfully"}
    json.NewEncoder(w).Encode(response)
}

func main() {
    http.HandleFunc("/ship-order", shipOrderHandler)
    http.HandleFunc("/cancel-shipping", cancelShippingHandler)

    log.Println("Shipping Service running on port 8083...")
    log.Fatal(http.ListenAndServe(":8083", nil))
}
