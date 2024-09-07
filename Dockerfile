FROM golang:1.22-alpine
LABEL authors="caiords"

WORKDIR /app

EXPOSE 80
EXPOSE 8080
EXPOSE 8000

ENV DATABASE_URL="postgresql://root:<SUA_SENHA>@<SUA_URL>:5432/fiscaliza?sslmode=disable

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main ./cmd

CMD ["./main"]