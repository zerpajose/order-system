package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"google.golang.org/grpc"
	pb "github.com/zerpajose/order-system/proto/orders"
)

type Order struct {
	OrderID    string   `json:"order_id"`
	ProductIDs []string `json:"product_ids"`
}

var orderClient pb.OrderServiceClient

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

	// Forward request to Orders service via gRPC
	resp, err := orderClient.CreateOrder(context.Background(), &pb.CreateOrderRequest{
		OrderId:    order.OrderID,
		ProductIds: order.ProductIDs,
	})
	if err != nil {
		http.Error(w, "Error creating order", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func main() {
	// Set up gRPC connection to orders service
	conn, err := grpc.Dial("orders-service:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect to orders service: %v", err)
	}
	defer conn.Close()
	orderClient = pb.NewOrderServiceClient(conn)

	// Set up HTTP server for external clients
	http.HandleFunc("/orders", handleNewOrder)
	log.Printf("API Gateway listening at :8090")
	if err := http.ListenAndServe(":8090", nil); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
