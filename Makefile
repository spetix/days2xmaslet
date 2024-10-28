ARGS?=

setup:
	go install github.com/boumenot/gocover-cobertura@latest
	go install github.com/gotesttools/gotestfmt/v2/cmd/gotestfmt@latest
	go mod download
	go mod tidy

linux-amd64: 
	@echo "Building linux-amd64"
	mkdir -p bin/linux-amd64
	GOOS=linux GOARCH=amd64 go build -v -o bin/linux-amd64/ ./cmd/...

darwin-amd64:
	@echo "Building darwin-amd64"
	mkdir -p bin/darwin-amd64
	GOOS=darwin GOARCH=amd64 go build -v -o bin/darwin-amd64/ ./cmd/...

windows-amd64:
	@echo "Building windows-amd64"
	mkdir -p bin/windows-amd64
	GOOS=windows GOARCH=amd64 go build -v -o bin/windows-amd64/ ./cmd/...

linux-arm64:
	@echo "Building linux-arm64"
	mkdir -p bin/linux-arm64
	GOOS=linux GOARCH=arm64 go build -v -o bin/linux-arm64/ ./cmd/...

# Define the build-all target
build-all: linux-amd64 darwin-amd64 windows-amd64 linux-arm64

# TODO clean binaries
clean:
	go clean ./cmd/...
	rm -Rf bin
	rm TestResult*.json
	rm test.log
	rm coverage.*

# TODO run tests
test:
	go test -race -json -v -coverprofile=coverage.txt ./... 2>&1 | tee test.log | gotestfmt

coverage:
	gocover-cobertura < coverage.txt > coverage.xml
