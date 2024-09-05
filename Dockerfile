FROM golang:1.22-alpine
LABEL authors="caiords"

WORKDIR /app

EXPOSE 80
EXPOSE 8080
EXPOSE 8000

ENV DATABASE_URL="postgresql://greentech:jmIBU2hg2Hp8jvzrVDjhFfGQepgCTNdb@dpg-crcu2tqj1k6c73ctl9qg-a.oregon-postgres.render.com/fiscaliza"

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main ./cmd

CMD ["./main"]