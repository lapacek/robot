build_dir = "build"
executable = "robot"
executable_arm = "robot.linux_arm"

clean_build_dir:
	@[ -d $(build_dir) ] && rm -rf $(build_dir) || exit 0

test-tracker:
	@go clean -testcache
	go mod tidy && \
	@cd cmd/robot/cmd/tracker && \
	go test -v ./... | grep -v "no test files"

.PHONY: build-ev3
build-ev3:
	@docker build --platform=linux/arm -t ev3arm32v5go -f ./script/Dockerfile . && \
	docker tag ev3arm32v5go ev3arm && \
	docker run --rm --platform linux/arm -v $(shell pwd)/$(build_dir):/$(build_dir) -it ev3arm /bin/sh -c "cp -r /app/build/$(executable_arm) /$(build_dir)/"

.PHONY: compile-ev3
compile-ev3:
	@if [ -f $(build_dir)/$(executable_arm) ]; then rm $(build_dir)/$(executable_arm); fi && \
	cd cmd/robot/ && \
	env GOOS=linux GOARCH=arm CGO_ENABLED=1 go build -o ./../../$(build_dir)/$(executable_arm)

.PHONY: test
test: test-tracker

.PHONY: clean
clean: clean_build_dir
	@if [ -f $(executable_name) ]; then rm $(executable_name); fi

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
	@echo "  install: build and install the project"
	@echo "  build: build the project"
	@echo "  build-ev3: build the project binary for the ev3 brick"
	@echo "  compile-ev3: cross compile the project for ev3 brick"
	@echo "  test: run tests"
	@echo "  clean: clean the project"

.PHONY: all
all: install

