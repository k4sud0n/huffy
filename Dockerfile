FROM golang:1.23-alpine AS builder

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

COPY go.mod ./
RUN go mod download

COPY . .

EXPOSE 3000

CMD ["air", "-c", ".air.toml"]