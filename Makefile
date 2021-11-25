test:

	@go clean -testcache

	@go test -v ./... | grep -v "no test files"

.PHONY: test
