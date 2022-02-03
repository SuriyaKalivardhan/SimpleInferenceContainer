init:
	go mod init SimpleInferencer

clean:
	rm -rf cmd/simpleinferencer/simpleinferencer.exe

build:
	go build -o cmd/simpleinferencer/simpleinferencer.exe cmd/simpleinferencer/main.go

run:
	go run cmd/simpleinferencer/main.go

req:
	curl -X POST localhost:5001/inference -H "Content-Type:application/json" -d '{"Id":"1234", "Type": "Infer", "Input":"Magi"}' | jq

all:
	init clean build run

docker-build:
	docker build --no-cache -f Dockerfile -t suriyakalivardhan/simpleinferencer:v0 .

docker-push:
	docker push suriyakalivardhan/simpleinferencer:v0
