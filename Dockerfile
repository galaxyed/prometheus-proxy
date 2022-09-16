FROM golang:alpine AS builder

WORKDIR /go/src/app

COPY go.mod ./

RUN go mod download

COPY . .

FROM alpine AS runner

WORKDIR /app

COPY --from=builder /go/src/app/prometheus-proxy /app/prometheus-proxy

COPY .conf.yml conf.yml

ENTRYPOINT [ "prometheus-proxy", "--config", "conf.yml" ]