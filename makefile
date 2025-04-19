BINARY_NAME=sharego
DIST_DIR=dist

.PHONY: all vet lint test build clean ci release run build-for-all tls 

all: vet lint test build

tls:
	@echo "Generating TLS certificates..."
	@mkdir -p ./certs
	@openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout ./certs/key.pem -out ./certs/cert.pem
	@echo "TLS certificates generated."

vet:
	@echo "Running go vet..."
	@go vet ./...
	@echo "Go vet completed."

lint:
	@echo "Running golangci-lint..."
	@golangci-lint run ./...
	@echo "Golangci-lint completed."

test:
	@echo "Running tests..."
	@go test -v ./... | tee result.log | tail -20
	@echo "Tests completed."
	@echo "Test results:"
	@tail -n 20 result.log
	@echo "Test results saved to result.log."

ci: 
	@echo "Running CI tasks..."
	vet lint test build-for-all
	@echo "CI tasks completed."

run:
	@go build -o ./bin/app
	@./bin/app


clean:
	@rm -rf ./bin
	@rm -rf $(DIST_DIR)
	@rm -rf ./certs
	@rm -f *.log
	@echo Cleaned up bin, dist, and certs directories.
	@echo Clean complete!


build:
	@go build -o ./bin/app
	@echo Build complete!

build-for-all:
	@GOOS=windows GOARCH=amd64 go build -o $(DIST_DIR)/$(BINARY_NAME)-windows-amd64.exe
	@GOOS=linux GOARCH=amd64 go build -o $(DIST_DIR)/$(BINARY_NAME)-linux-amd64
	@GOOS=linux GOARCH=arm go build -o $(DIST_DIR)/$(BINARY_NAME)-linux-arm
	@GOOS=linux GOARCH=arm64 go build -o $(DIST_DIR)/$(BINARY_NAME)-linux-arm64
	@GOOS=darwin GOARCH=amd64 go build -o $(DIST_DIR)/$(BINARY_NAME)-darwin-amd64
	@GOOS=darwin GOARCH=arm64  go build -o $(DIST_DIR)/$(BINARY_NAME)-darwin-arm64
	@echo Build complete!