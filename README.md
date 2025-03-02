# Order System Microservices

This project implements a microservices-based order system with two services:
- Products Service: Manages product information (internal gRPC service)
- Orders Service: Handles order creation and management (external HTTP API)

## Architecture

- Products Service runs on port 50051 (gRPC)
- Orders Service runs on port 8080 (HTTP)
- Communication between services is done via gRPC
- Both services are containerized and ready for Kubernetes deployment

## Development

1. Generate gRPC code (requires protoc):
```bash
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative products.proto
```

2. Build Docker images:
```bash
docker build -t products-service:latest ./products-service
docker build -t orders-service:latest ./orders-service
```

3. Deploy to Kubernetes:
```bash
kubectl apply -f k8s/products-deployment.yaml
kubectl apply -f k8s/orders-deployment.yaml
```

## API Usage

Create a new order:
```bash
curl -X POST http://[ORDERS_SERVICE_IP]/orders \
  -H "Content-Type: application/json" \
  -d '{"order_id": "123", "product_ids": ["1", "2"]}'
```
