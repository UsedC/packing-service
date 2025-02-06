FROM golang:1.23.2-alpine as builder

WORKDIR /app

COPY . .

RUN go build -o main ./cmd/main/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .

COPY .env .

EXPOSE 8080

CMD ["./main"]