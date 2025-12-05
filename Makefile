qa: analyze test

analyze:
	@go vet ./...

test:
	@go test -cover ./...

clean:
	@rm -rf build/

build-ui:
	@cd ui && npm run build

build: clean build-ui
	@GOOS=darwin GOARCH=amd64 go build -o ./build/techlounge-darwin-amd64 main.go
	@GOOS=darwin GOARCH=arm64 go build -o ./build/techlounge-darwin-arm64 main.go
	@GOOS=linux GOARCH=amd64 go build -o ./build/techlounge-linux-amd64 main.go
	@GOOS=linux GOARCH=arm64 go build -o ./build/techlounge-linux-arm64 main.go
	@GOOS=windows GOARCH=amd64 go build -o ./build/techlounge-windows-amd64.exe main.go
	@GOOS=windows GOARCH=arm64 go build -o ./build/techlounge-windows-arm64.exe main.go

build-docker:
	@docker build -t thenativeweb/techlounge .

.PHONY: analyze \
				build \
				build-docker \
				build-ui \
				clean \
				qa \
				test
