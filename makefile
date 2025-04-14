run:
	@go build -o ./bin/app
	@./bin/app
clean:
	@rm -r ./bin
build:
	@go build -o ./bin/app.exe
	@echo Build complete!

build-for-all:
	@GOOS=windows GOARCH=amd64 go build -o ./bin/win/app.exe
	@GOOS=linux GOARCH=amd64 go build -o ./bin/linux/amd/app
	@GOOS=linux GOARCH=arm go build -o ./bin/linux/arm/app
	@GOOS=darwin GOARCH=amd64 go build -o ./bin/mac/amd/app
	@GOOS=darwin GOARCH=arm64  go build -o ./bin/mac/arm/app
	@echo Build complete!