BINARY_NAME = go_project_euler

build:
	go build .

problem ?= 1

run: build
	./${BINARY_NAME} ${problem} ${input}

clean:
	go clean

test:
	go test -v -cover ./...

fmt:
	go fmt ./...

vet:
	go vet ./...