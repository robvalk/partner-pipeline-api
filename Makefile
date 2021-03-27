.default: clean build

clean:
	rm -rf ./bin

build:
	go build -o bin/api-developer-v1 cmd/api-developer-v1/main.go

test:
	go test ./pkg/*

run:
	go run ./cmd/api-developer-v1/main.go

# Helpers

swaggerGen:
	oapi-codegen -generate chi-server spec/api-definition.yaml > handlers/handlers.go
	oapi-codegen -generate types spec/api-definition.yaml > models/models.go
	oapi-codegen -generate spec spec/api-definition.yaml > spec/spec.go