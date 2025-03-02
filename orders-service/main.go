package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"google.golang.org/grpc"
	pb "order-system/proto"
)

type Order struct {
	OrderID    string   `json:"order_id"`
	ProductIDs []string `json:"product_ids"`
}

var productClient pb.ProductServiceClient

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

	// Get product details for each product in the order
	var products []*pb.ProductResponse
	for _, productID := range order.ProductIDs {
		product, err := productClient.GetProduct(context.Background(), &pb.ProductRequest{ProductId: productID})
		if err != nil {
			http.Error(w, "Error getting product details", http.StatusInternalServerError)
			return
		}
		products = append(products, product)
	}

	// TODO: Save order to database

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"order_id": order.OrderID,
		"products": products,
	})
}

func main() {
	// Set up gRPC connection to products service
	conn, err := grpc.Dial("products-service:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect to products service: %v", err)
	}
	defer conn.Close()
	productClient = pb.NewProductServiceClient(conn)

	// Set up HTTP server
	http.HandleFunc("/orders", handleNewOrder)
	log.Printf("Orders service listening at :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
