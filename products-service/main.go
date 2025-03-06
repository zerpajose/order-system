package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Product struct {
	ProductID string  `json:"product_id"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	Stock     int32   `json:"stock"`
}

var products = map[string]Product{
	"1": {ProductID: "1", Name: "Laptop", Price: 999.99, Stock: 50},
	"2": {ProductID: "2", Name: "Smartphone", Price: 599.99, Stock: 100},
}

func getProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.URL.Query().Get("product_id")
	if product, ok := products[productID]; ok {
		json.NewEncoder(w).Encode(product)
		return
	}
	http.Error(w, "Product not found", http.StatusNotFound)
}

func main() {
	http.HandleFunc("/product", getProduct)
	log.Println("Products service listening on port 50051")
	log.Fatal(http.ListenAndServe(":50051", nil))
}
