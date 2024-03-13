FROM golang:1.22-alpine

WORKDIR /app

COPY . .

run go mod tidy
run go build -o main cmd/main.go

EXPOSE 3000

# CMD ["./main"]
CMD ["./main"]


# RUN go mod download

COPY cmd/main.go /
COPY go.mod .

# RUN go build -o cmd/main
# RUN go mod tidy .

# run go build -o bin/main ./cmd
# # RUN go run pagetest/cmd/main.go
# RUN go mod tidy
#
# CMD ["./bin./main"]
#
# COPY . .
#
# RUN go build -o cmd/main.go
#
# # EXPOSE 8000

