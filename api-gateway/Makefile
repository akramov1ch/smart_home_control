build:
	docker build -t api-gateway .

run:
	docker run -p 9090:9090 -p 8080:8080 --env-file .env api-gateway
