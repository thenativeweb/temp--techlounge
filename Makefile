qa: analyze test

analyze:
	@go vet ./...

test:
	@go test -cover ./...

build-docker:
	@docker build -t thenativeweb/techlounge .

.PHONY: analyze \
				build-docker \
				qa \
				test
