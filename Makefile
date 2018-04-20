all: build

build: build-config build-proxy

build-config:
	dep ensure
	go build -o bin/config ./cmd/config

build-proxy:
	dep ensure
	go build -o bin/proxy ./cmd/proxy