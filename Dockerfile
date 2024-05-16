FROM golang:latest
WORKDIR /app
COPY . .
RUN go build -o logger cmd/main.go
CMD ["./logger", "test_file.txt"]