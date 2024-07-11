FROM golang:latest AS build

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o main ./cmd/main.go

EXPOSE 8000

CMD ["/app/main"]
