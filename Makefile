BINARY_NAME = go_project_euler

build:
	go build .

problem ?= 1
input ?= 0

run: build
	./${BINARY_NAME} --problem ${problem} --input ${input}

clean:
	go clean

test:
	go test -v -coverprofile=coverage.out ./...

fmt:
	go fmt ./...

vet:
	go vet ./...