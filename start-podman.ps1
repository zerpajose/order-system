# Build images
Write-Host "Building images..."
podman build -t products-service:v1.0.0 ./products-service
podman build -t orders-service:v1.0.0 ./orders-service
podman build -t api-gateway:v1.0.0 ./api-gateway

# Clean up existing pod if it exists
Write-Host "Cleaning up existing pod..."
podman pod rm -f order-system-pod

# Create pod
Write-Host "Creating pod..."
podman pod create --name order-system-pod -p 8090:8090 -p 50051:50051 -p 50052:50052

# Run services
Write-Host "Starting services..."
podman run -d --pod order-system-pod --name products-service products-service:v1.0.0
podman run -d --pod order-system-pod --name orders-service orders-service:v1.0.0
podman run -d --pod order-system-pod --name api-gateway api-gateway:v1.0.0

Write-Host "System is running!"
Write-Host "API Gateway is accessible at http://localhost:8090"
