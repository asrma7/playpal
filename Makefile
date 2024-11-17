run-all:
	cd api-gateway && start /b make server
	cd auth-svc && start /b make server
	cd feed-svc && start /b make server

view-ports:
	netstat -ano | findstr :3000
	netstat -ano | findstr :50051
	netstat -ano | findstr :50052

stop:
	# Kill the process running on port 3000
	taskkill /F /PID $(shell netstat -ano | findstr :3000 | awk '{print $$5}') || true
	# Kill the process running on port 50051
	taskkill /F /PID $(shell netstat -ano | findstr :50051 | awk '{print $$5}') || true
	# Kill the process running on port 50052
	taskkill /F /PID $(shell netstat -ano | findstr :50052 | awk '{print $$5}') || true

docker-build:
	cd api-gateway && docker build -t ghcr.io/asrma7/playpal-api-gateway .
	cd auth-svc && docker build -t ghcr.io/asrma7/playpal-auth-svc .
	cd feed-svc && docker build -t ghcr.io/asrma7/playpal-feed-svc .

docker-push:
	docker push ghcr.io/asrma7/playpal-api-gateway
	docker push ghcr.io/asrma7/playpal-auth-svc
	docker push ghcr.io/asrma7/playpal-feed-svc

docker-run:
	docker docker compose up
