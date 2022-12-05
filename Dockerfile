FROM docker.io/golang:1.19.3-alpine3.17

WORKDIR /gambler

COPY . /gambler

RUN go build -o \
    ./bin/gambler-market-sync \
    ./src/cmd/marketdata/main.go
