# Stage 1 - build
FROM golang:1.24.2-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o admin_service cmd/admin_service/main.go

# Stage 2 - run
FROM alpine
WORKDIR /app
COPY --from=builder /app/admin_service .
CMD ["./admin_service"]