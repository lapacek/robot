build_dir = "build"
executable_name = "main"

clean_build_dir:
	@[ -d $(build_dir) ] && rm -rf $(build_dir) || exit 0

clean: clean_build_dir

build: clean
	@mkdir $(build_dir)
	@cd ./cmd && \
	go mod tidy && \
 	go build -o ../build/$(executable_name) main.go

test:
	@go clean -testcache
	@go test -v ./... | grep -v "no test files"

.PHONY: clean build test
