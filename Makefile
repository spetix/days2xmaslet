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

bin/linux-arm64:
	@echo "Building linux-arm64"
	mkdir -p bin/linux-arm64
	GOOS=linux GOARCH=arm64 go build -v -o bin/linux-arm64/ ./cmd/...



# Define the build-all target
.PHONY: build
build: 
	$(MAKE) linux-amd64 
	$(MAKE) darwin-amd64 
	$(MAKE) windows-amd64 
	$(MAKE) bin/linux-arm64


# TODO clean binaries
clean:
	go clean ./cmd/...
	rm -Rf bin
	rm -f test-result*.json
	rm -f coverage.*

test: build
	$(MAKE) test-results.json

# TODO run tests
test-results.json:
	go test -race -json -v -coverprofile=coverage.txt ./... 2>&1 | tee test-results.json | gotestfmt

coverage: test-results.json
	gocover-cobertura < coverage.txt > coverage.xml
