FROM golang:alpine AS builder

WORKDIR /go/src/app

COPY go.mod ./

RUN go mod download

COPY . .

RUN go build

FROM alpine AS runner

WORKDIR /app

COPY --from=builder /go/src/app/prometheus-proxy /app/prometheus-proxy

COPY conf.yml conf.yml

ENTRYPOINT [ "/app/prometheus-proxy", "--config", "/app/conf.yml" ]
