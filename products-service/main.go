package main

import (
	"context"
	"log"
	"net"
	"google.golang.org/grpc"
	pb "order-system/proto"
)

type server struct {
	pb.UnimplementedProductServiceServer
}

func (s *server) GetProduct(ctx context.Context, req *pb.ProductRequest) (*pb.ProductResponse, error) {
	// TODO: Implement actual product retrieval from database
	return &pb.ProductResponse{
		ProductId: req.ProductId,
		Name:      "Sample Product",
		Price:     99.99,
		Stock:     100,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterProductServiceServer(s, &server{})
	log.Printf("Products service listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
