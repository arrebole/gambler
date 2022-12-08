
all: gambler gamblerDataSync

gambler: src/cmd/gambler/main.go
	go build -o ./bin/gambler ./src/cmd/gambler/main.go

gamblerDataSync: src/cmd/marketdata/main.go
	go build -o ./bin/gambler-data-sync ./src/cmd/marketdata/main.go
