# Test creating an order
$orderData = @{
    order_id = "123"
    product_ids = @("1", "2")
} | ConvertTo-Json

Write-Host "Creating order..."
$response = Invoke-RestMethod -Uri "http://localhost:8090/orders" -Method Post -Body $orderData -ContentType "application/json"
Write-Host "Response:"
$response | ConvertTo-Json

# View pod status
Write-Host "`nPod Status:"
podman pod ps

# View container logs
Write-Host "`nAPI Gateway logs:"
podman logs api-gateway
Write-Host "`nOrders Service logs:"
podman logs orders-service
Write-Host "`nProducts Service logs:"
podman logs products-service
