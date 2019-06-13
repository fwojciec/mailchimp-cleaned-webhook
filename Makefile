.PHONY: build clean deploy

build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/cleaned-email-notify cleaned-email-notify/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/ping ping/main.go

clean:
	rm -rf ./bin

deploy: clean build
	sls deploy --verbose
