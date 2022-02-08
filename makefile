init:
	go mod init SimpleInferencer

clean:
	rm -rf cmd/simpleinferencer/simpleinferencer.exe

build:
	go build -o cmd/simpleinferencer/simpleinferencer.exe cmd/simpleinferencer/main.go

run:
	go run cmd/simpleinferencer/main.go

req:
	curl -X POST localhost:5005/inference -H "Content-Type:application/json" -d '{"Id":"1234", "Type": "Infer", "Input":"Magi"}' | jq

all:
	init clean build run

docker-build:
	docker build --no-cache -f Dockerfile -t suriyakeuapregistry.azurecr.io/suriyakalivardhan/simpleinferencer:v1 .

docker-push:
	docker push suriyakeuapregistry.azurecr.io/suriyakalivardhan/simpleinferencer:v1
