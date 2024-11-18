.PHONY: run-all view-ports stop docker-build docker-push docker-run

run-all:
ifeq ($(OS),Windows_NT)
	cd api-gateway && start /B make server
	cd auth-svc && start /B make server
	cd feed-svc && start /B make server
else
	cd api-gateway && make server &
	cd auth-svc && make server &
	cd feed-svc && make server &
endif

view-ports:
ifeq ($(OS),Windows_NT)
	netstat -ano | findstr :3000
	netstat -ano | findstr :50051
	netstat -ano | findstr :50052
else
	ss -tuln | grep :3000 || true
	ss -tuln | grep :50051 || true
	ss -tuln | grep :50052 || true
endif

stop:
ifeq ($(OS),Windows_NT)
	-for /f "tokens=5" %a in ('netstat -ano ^| findstr :3000') do taskkill /F /PID %a
	-for /f "tokens=5" %a in ('netstat -ano ^| findstr :50051') do taskkill /F /PID %a
	-for /f "tokens=5" %a in ('netstat -ano ^| findstr :50052') do taskkill /F /PID %a
else
	kill -9 $(shell lsof -ti:3000) || true
	kill -9 $(shell lsof -ti:50051) || true
	kill -9 $(shell lsof -ti:50052) || true
endif

docker-build:
	cd api-gateway && docker build -t ghcr.io/asrma7/playpal-api-gateway .
	cd auth-svc && docker build -t ghcr.io/asrma7/playpal-auth-svc .
	cd feed-svc && docker build -t ghcr.io/asrma7/playpal-feed-svc .

docker-push:
	docker push ghcr.io/asrma7/playpal-api-gateway
	docker push ghcr.io/asrma7/playpal-auth-svc
	docker push ghcr.io/asrma7/playpal-feed-svc

docker-run:
	docker compose up

