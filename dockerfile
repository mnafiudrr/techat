# Gunakan image Go dengan Alpine
FROM golang:1.21-alpine AS builder

# Set working directory
WORKDIR /app

# Install git (diperlukan untuk go mod download)
RUN apk add --no-cache git

# Copy semua file
COPY . .

# Download dependencies
RUN go mod tidy

# Build dengan CGO disabled agar tidak bergantung pada GLIBC
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app .

# Stage kedua: gunakan Alpine yang lebih kecil
FROM alpine:latest

# Set working directory
WORKDIR /root/

# Copy binary dari stage sebelumnya
COPY --from=builder /app/app .

# Copy file .env
COPY .env .

# Expose port API
EXPOSE ${EXPOSED_PORT}

# Jalankan aplikasi
CMD ["./app"]
