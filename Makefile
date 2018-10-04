BINARY_NAME=microdungeon

build:
	go build -o $(BINARY_NAME) -v

test:
	go test  -v ./...

clean:
	rm -rf $(BINARY_NAME)

.DEFAULT_GOAL := test
