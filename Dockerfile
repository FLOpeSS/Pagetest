FROM golang:1.22-alpine

WORKDIR /app

COPY . .

run go mod tidy
run go build -o main cmd/main.go

EXPOSE 3000

CMD ["./main"]



COPY cmd/main.go /
COPY go.mod .
