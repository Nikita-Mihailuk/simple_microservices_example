# Stage 1 - build
FROM golang:1.24.2-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o user_service cmd/user_service/main.go
RUN go build -o migrator cmd/migrator/main.go

# Stage 2 - run
FROM alpine
WORKDIR /app
COPY --from=builder /app/user_service .
COPY --from=builder /app/migrator .
COPY ./migrations ./migrations
CMD ["./user_service"]