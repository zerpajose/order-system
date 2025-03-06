package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Product struct {
	ProductID string  `json:"product_id"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	Stock     int32   `json:"stock"`
}

type Order struct {
	OrderID    string    `json:"order_id"`
	ProductIDs []string  `json:"product_ids"`
	Products   []Product `json:"products"`
}

func handleNewOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var order Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Forward request to Orders service via HTTP
	orderData, err := json.Marshal(order)
	if err != nil {
		http.Error(w, "Error in JSON encoding", http.StatusInternalServerError)
		return
	}

	resp, err := http.Post("http://orders-service:50052/orders", "application/json", bytes.NewBuffer(orderData))
	if err != nil {
		log.Printf("Error making request to Orders service: %v", err)
		http.Error(w, "Error fetching the service", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		log.Printf("Orders service responded with status %d: %s", resp.StatusCode, string(body))
		http.Error(w, "Error creating order", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}

func main() {
	http.HandleFunc("/orders", handleNewOrder)
	log.Println("API Gateway listening on port 8090")
	log.Fatal(http.ListenAndServe(":8090", nil))
}
