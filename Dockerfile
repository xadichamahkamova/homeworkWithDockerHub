FROM golang:1.22.3-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o main cmd/main.go

FROM alpine:3.16

WORKDIR /app

COPY --from=builder /app/main .
COPY .env .

CMD ["/app/main"]
