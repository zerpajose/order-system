package main

import (
	"context"
	"log"
	"net"
	"google.golang.org/grpc"
	pb "github.com/zerpajose/order-system/proto/orders"
	productpb "github.com/zerpajose/order-system/proto/products"
)

type server struct {
	pb.UnimplementedOrderServiceServer
	productClient productpb.ProductServiceClient
}

func (s *server) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	// Get product details for each product in the order
	var products []*productpb.ProductResponse
	for _, productID := range req.ProductIds {
		product, err := s.productClient.GetProduct(ctx, &productpb.ProductRequest{ProductId: productID})
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	// TODO: Save order to database

	return &pb.CreateOrderResponse{
		OrderId:  req.OrderId,
		Products: products,
	}, nil
}

func main() {
	// Set up gRPC connection to products service
	productConn, err := grpc.Dial("products-service:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect to products service: %v", err)
	}
	defer productConn.Close()

	// Create the server instance with product client
	s := &server{
		productClient: productpb.NewProductServiceClient(productConn),
	}

	// Start the gRPC server
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterOrderServiceServer(grpcServer, s)

	log.Printf("Orders service listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
