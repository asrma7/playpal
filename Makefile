run-all:
	cd api-gateway && start /b make server
	cd auth-service && start /b make server

view-ports:
	netstat -ano | findstr :3000
	netstat -ano | findstr :50051

stop:
	# Kill the process running on port 3000
taskkill /F /PID $(shell netstat -ano | findstr :3000 | awk '{print $$5}') || true
	# Kill the process running on port 50051
	taskkill /F /PID $(shell netstat -ano | findstr :50051 | awk '{print $$5}') || true

docker-build:
	cd api-gateway && docker build -t ghcr.io/asrma7/playpal-api-gateway .
	cd auth-service && docker build -t ghcr.io/asrma7/playpal-auth-service .

docker-push:
	docker push ghcr.io/asrma7/playpal-api-gateway
	docker push ghcr.io/asrma7/playpal-auth-service

docker-run:
	docker docker compose up
