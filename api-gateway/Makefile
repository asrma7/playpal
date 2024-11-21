proto:
	protoc --go_out=. --go-grpc_out=. pkg/proto/auth/pb/*.proto
	protoc --go_out=. --go-grpc_out=. pkg/proto/feed/pb/*.proto


server:
	go run cmd/main.go
