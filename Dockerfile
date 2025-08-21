FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod tidy && go build -o app ./cmd/app

FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/app .
CMD ["./app"]
