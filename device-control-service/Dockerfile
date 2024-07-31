FROM golang:1.22.5 AS builder

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o main cmd/main.go

FROM alpine:latest

COPY --from=builder /app/main /main

EXPOSE 50051

CMD ["./main"]
