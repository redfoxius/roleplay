FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o chat-service

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/chat-service .

EXPOSE 8082

CMD ["./chat-service"] 