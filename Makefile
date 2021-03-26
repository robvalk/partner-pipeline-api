.default: clean build

clean:
	rm -rf ./bin

build:
	go build -o bin/api-developer-v1 cmd/api-developer-v1/main.go

run:
	go run cmd/api-developer-v1/main.go