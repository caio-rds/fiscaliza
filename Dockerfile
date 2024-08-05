FROM golang:1.20-alpine
LABEL authors="caiords"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main ./cmd

CMD ["./main"]