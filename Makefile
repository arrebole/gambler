
all: gamblerMarketSync

gamblerMarketSync: src/cmd/marketdata/main.go
	go build -o ./bin/gambler-market-sync ./src/cmd/marketdata/main.go
