# Build images from root directory to include proto
Write-Host "Building images..."
docker build -t products-service:v1.0.0 -f ./products-service/Dockerfile .
docker build -t orders-service:v1.0.0 -f ./orders-service/Dockerfile .
docker build -t api-gateway:v1.0.0 -f ./api-gateway/Dockerfile .

# Create pod
Write-Host "Creating pod..."
docker pod rm -f order-system-pod
docker pod create --name order-system-pod -p 8090:8090 -p 50051:50051 -p 50052:50052

# Run services
Write-Host "Starting services..."
docker run -d --pod order-system-pod --name products-service products-service:v1.0.0
docker run -d --pod order-system-pod --name orders-service orders-service:v1.0.0
docker run -d --pod order-system-pod --name api-gateway api-gateway:v1.0.0

Write-Host "System is running!"
Write-Host "API Gateway is accessible at http://localhost:8090"
