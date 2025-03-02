package main

import (
	"context"
	"log"
	"net"
	"google.golang.org/grpc"
	pb "github.com/zerpajose/order-system/proto/products"
)

type server struct {
	pb.UnimplementedProductServiceServer
	products map[string]*pb.ProductResponse
}

func (s *server) GetProduct(ctx context.Context, req *pb.ProductRequest) (*pb.ProductResponse, error) {
	if product, ok := s.products[req.ProductId]; ok {
		return product, nil
	}
	return nil, nil // Return nil if product not found
}

func main() {
	// Initialize server with some sample products
	s := &server{
		products: map[string]*pb.ProductResponse{
			"1": {
				ProductId: "1",
				Name:     "Laptop",
				Price:    999.99,
				Stock:    50,
			},
			"2": {
				ProductId: "2",
				Name:     "Smartphone",
				Price:    599.99,
				Stock:    100,
			},
		},
	}

	// Start gRPC server
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterProductServiceServer(grpcServer, s)

	log.Printf("Products service listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
