.PHONY: build clean deploy

build:
	rm -f bin/*
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/aws ./endpoints/main_lambda.go

clean:
	rm -rf ./bin

deploy: clean build
	sls deploy --verbose