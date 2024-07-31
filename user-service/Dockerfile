FROM golang:1.20-alpine

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o user-service ./cmd/main.go

EXPOSE 8081

CMD ["./user-service"]
