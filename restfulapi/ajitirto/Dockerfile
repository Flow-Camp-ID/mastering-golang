# Tahap 1: Bangun aplikasi Go
FROM golang:1.23 as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

# Tahap 2: Buat image runtime minimal
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]