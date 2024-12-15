FROM golang:1.23.1 AS builder

WORKDIR /app

COPY ./go.mod ./

COPY ./go.sum ./

RUN go mod download

COPY . .

RUN go build -o main ./cmd/main.go

FROM ubuntu:latest

WORKDIR /app

COPY --from=builder /app/main .

CMD ["./main"]