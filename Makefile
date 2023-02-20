build_dir = "build"
executable_name = "robot"

clean_build_dir:
	@[ -d $(build_dir) ] && rm -rf $(build_dir) || exit 0

test-tracker:
	@go clean -testcache
	@cd cmd/robot/cmd/tracker && \
	go mod tidy && \
	go test -v ./... | grep -v "no test files"

.PHONY: test
test: test-tracker

.PHONY: clean
clean: clean_build_dir
	@rm -rf $(executable_name)

.PHONY: build
build: clean
	@mkdir $(build_dir) && \
	cd cmd/robot/ && \
	go mod tidy && \
	go build -o ../../$(build_dir)/$(executable_name)

.PHONY: install
install: build
	@ln -s $(PWD)/$(build_dir)/$(executable_name) $(executable_name)

.PHONY: help
help:
	@echo "Usage: make [target]"
	@echo "Targets:"
	@echo "  build: build the project"
	@echo "  install: build and install the project"
	@echo "  test: run tests"
	@echo "  clean: clean the project"

.PHONY: all
all: install

