build_dir = "build"
executable_name = "main"

clean_build_dir:
	@[ -d $(build_dir) ] && rm -rf $(build_dir) || exit 0

.PHONY: clean
clean: clean_build_dir

.PHONY: build
build: clean
	@mkdir $(build_dir)
	@cd ./cmd && \
	go mod tidy && \
 	go build -o ../build/$(executable_name) main.go

.PHONY: test
test:
	@go clean -testcache
	@cd internal && \
	go test -v ./... | grep -v "no test files"
