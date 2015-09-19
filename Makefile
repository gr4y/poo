default: clean test build

clean:
	rm -rf build

test:
	go test -v ./...

build:
	go build
