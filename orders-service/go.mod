module orders-service

go 1.21

require (
	github.com/zerpajose/order-system/proto v0.0.0
	google.golang.org/grpc v1.64.0
	google.golang.org/protobuf v1.36.5
)

require (
	golang.org/x/net v0.32.0 // indirect
	golang.org/x/sys v0.28.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20241202173237-19429a94021a // indirect
)

replace github.com/zerpajose/order-system/proto => ../proto
