.PHONY: build clean deploy

build:
	rm -f bin/*
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/aws main.go

clean:
	rm -rf ./bin

deploy: clean build
	sls deploy --verbose