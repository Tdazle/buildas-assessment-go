# Dockerfile.backend
FROM golang:1.20-alpine
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o myapp ./cmd/myapp/main.go
CMD ["./myapp"]
