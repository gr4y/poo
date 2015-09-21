default: clean test build

coverage:
	go test -coverprofile=coverage.out -coverpkg ./...
	go tool cover -func=coverage.out
	go tool cover -html=coverage.out -o coverage.html

clean:
	go clean
	rm -rf build

test:
	go test -v ./...

build:
	go build
