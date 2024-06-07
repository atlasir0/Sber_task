FROM golang:1.21.1-alpine AS builder


WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main ./cmd/main.go

FROM golang:1.21.1-alpine

WORKDIR /app

COPY --from=builder /app/main .

COPY config.yaml ./

EXPOSE 8080

CMD ["./main"]