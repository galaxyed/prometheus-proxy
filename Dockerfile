FROM golang:alpine AS builder

WORKDIR /go/src/app

COPY go.mod ./

RUN go mod download

COPY . .

CMD ["go", "run", "main.go"]