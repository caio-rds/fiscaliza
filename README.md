# O que é este projeto?
> Esse projeto é um portal onde os moradores podem se cadastrar e informar sobre crimes que ocorreram em sua região,
> com o intuito de alertar a população sobre os locais mais perigosos e diminuir a criminalidade.


# Como surgiu a ideia?
> Após entrevistar moradores da vizinhança e comerciantes, percebemos que a criminalidade é um problema recorrente
> e isso causa medo e insegurança na população. Com isso, surgiu a ideia de criar um portal onde os moradores podem
> compartilhar informações sobre crimes que ocorreram em sua região.


# Qual a motivação para este projeto?
> A motivação para este projeto é a segurança da população, pois com a ajuda dos moradores, a polícia pode se antecipar
> e outros moradores podem evitar situações perigosas pois o compartilhamento é em tempo real.

# Como rodar o projeto?
> Instale Go, Docker e Docker Compose ou PostGres em sua máquina.

### Há algumas maneiras de fazer isso

#### 1. Docker / Docker Compose
### Dockerfile
```dockerfile
FROM golang:1.22-alpine
LABEL authors="caiords"

WORKDIR /app

EXPOSE 80
EXPOSE 8080
EXPOSE 8000

ENV DATABASE_URL="postgresql://root:<SUA_SENHA>@<SUA_URL>:5432/fiscaliza?sslmode=disable"

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main ./cmd

CMD ["./main"]
```

### docker-compose.yaml
```yaml
services:
  db:
    image: postgres:13
    container_name: db
    restart: always
    environment:
      POSTGRES_USER: root
      POSTGRES_PASSWORD: SUA_SENHA
      POSTGRES_DB: fiscaliza
    ports:
      - "5432:5432"
    volumes:
      - ./db:/var/lib/postgresql/data

  app:
    build:
        context: .
        dockerfile: Dockerfile
    container_name: fiscaliza_app
    ports:
      - "8000:8000"
    depends_on:
      - db
    environment:
      - DATABASE_URL=postgres://root:<SUA_SENHA>@db:5432/fiscaliza?sslmode=disable

    command: ["./main"]

volumes:
    db:
```
>1.1 Clone o repositório e execute o comando `docker-compose up` na raiz do projeto.
> 
>1.2 Clone o repositório e execute o comando `docker build -t fiscaliza .` e depois `docker run -p 8080:8080 fiscaliza`.

#### 2. Local
>2.1 Clone o repositório e execute o comando `go run cmd/main.go` na raiz do projeto.
> 
>2.2 Clone o repositório e execute o comando `go build -o fiscaliza cmd/main.go` e depois `./fiscaliza`.
> 
>2.3 Instale o PostGres e crie um banco de dados com o nome `fiscaliza`.

### Como criar as tabelas no banco de dados?
> O projeto já está configurado para executar uma migrations caso o banco de dados esteja vazio.
> 
> Se preocupe apenas em criar o banco de dados e o restante será feito automaticamente.

### Não esqueça de configurar o arquivo `.env` com as variáveis de ambiente.
```env
DATABASE_URL=<SUA_URL_DE_CONEXÃO>
SG_KEY=<SUA_API_KEY>
SG_TEMPLATE_ID=<SEU_TEMPLATE_ID>
TWILIO_PHONE=<SEU_TELEFONE>
TWILIO_SID=<SEU_SID>
TWILIO_TOKEN=<SEU_TOKEN>
```

# Libs utilizadas no projeto
[Gin](https://gin-gonic.com/docs/), [Gorm](https://gorm.io/docs/index.html), [JWT](https://pkg.go.dev/github.com/golang-jwt/jwt/v5), [Bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt), [CORS](https://pkg.go.dev/github.com/gin-contrib/cors), 
[PostGres Driver (PGX)](https://pkg.go.dev/github.com/jackc/pgx/v4)