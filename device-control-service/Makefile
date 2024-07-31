run:
	go run cmd/main.go

build:
	go build -o main cmd/main.go

docker-build:
	docker build -t device-control-service .

docker-run:
	docker run -p 50051:50051 device-control-service
