FROM golang:alpine AS builder

WORKDIR /go/src/app

COPY go.mod ./

RUN go mod download

COPY . .

RUN mkdir -p bin/ && go build -o ./bin/ ./...

FROM alpine AS runner

WORKDIR /

COPY --from=builder /go/src/app/bin/prometheus-proxy /prometheus-proxy

COPY conf.yml conf.yml

ENTRYPOINT [ "/prometheus-proxy", "--config", "conf.yml" ]
