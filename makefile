init:
	go mod init SimpleInferencer

clean:
	rm -rf cmd/simpleinferencer/simpleinferencer

build:
	go build -o cmd/simpleinferencer/simpleinferencer cmd/simpleinferencer/main.go

run:
	go run cmd/simpleinferencer/main.go

req:
	curl -X POST localhost:5001/inference -H "Content-Type:application/json" -d '{"Id":"Hi", "Type": "Infer", "Input":"Agazhi"}' | jq

all:
	init clean build run
