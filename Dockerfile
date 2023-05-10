FROM golang:latest

WORKDIR /app

COPY . .
RUN go build -o main ./cmd/apiserver/main.go
CMD ["/app/main"]
