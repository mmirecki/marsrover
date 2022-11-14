.PHONY: build \
	test \
	run \

build:
	go build -o MarsRover main.go

test:
	ginkgo ./...

run:
	go run main.go
