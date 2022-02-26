.PHONY: build clean deploy

build:
	dep ensure -v
	env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -v -ldflags "-d -s -w" -a -tags netgo -installsuffix netgo -o bin/ping ping/main.go
	env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -v -ldflags "-d -s -w" -a -tags netgo -installsuffix netgo -o bin/preview preview/main.go

clean:
	rm -rf ./bin ./vendor Gopkg.lock

deploy: clean build
	sls deploy --verbose
