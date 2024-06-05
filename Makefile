run:
	go run .

build:
	go build -o install

format:
	gofmt -e -d -s -w .
