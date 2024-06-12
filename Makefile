run:
	go run .

build: build-amd64 build-arm64

build-amd64:
	CGO_ENABLES=0 GOOS=linux GOARCH=amd64 go build -o amd-install

build-arm64:
	CGO_ENABLES=0 GOOS=linux GOARCH=arm64 go build -o arm-install

format:
	gofmt -e -d -s -w .
