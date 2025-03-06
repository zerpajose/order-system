package main

import (
	"encoding/json"
	"fmt"
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

func createOrder(w http.ResponseWriter, r *http.Request) {
	var order Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for _, productID := range order.ProductIDs {
		resp, err := http.Get(fmt.Sprintf("http://products-service:50051/product?product_id=%s", productID))
		if err != nil || resp.StatusCode != http.StatusOK {
			http.Error(w, "Error fetching product details", http.StatusInternalServerError)
			return
		}
		var product Product
		if err := json.NewDecoder(resp.Body).Decode(&product); err != nil {
			http.Error(w, "Error decoding product details", http.StatusInternalServerError)
			return
		}
		order.Products = append(order.Products, product)
	}

	// TODO: Save order to database

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}

func main() {
	http.HandleFunc("/orders", createOrder)
	log.Println("Orders service listening on port 50052")
	log.Fatal(http.ListenAndServe(":50052", nil))
}
