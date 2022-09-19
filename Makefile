GOHOSTOS:=$(shell go env GOHOSTOS)
GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)


.PHONY: all
all: run

.PHONY: build
# build
build:
	mkdir -p bin/ && go build -ldflags "-X main.Version=$(VERSION)" -o ./bin/ ./...


.PHONY: run
# start server
run: build
	./bin/prometheus-proxy --config conf.yml


.PHONY: docker-build
# start server
docker-build:
	docker build . -t prometheus-proxy:dev


.PHONY: docker-run
# start server
docker-run: docker-build
	docker run -it -p 8000:8000 prometheus-proxy:dev