
BINARY_NAME=todo_list
BUILD_DIR=bin
MAIN_PATH=cmd/main.go

BASE_DIR=$(shell pwd)

build:
	go build -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_PATH)

run:
	go run $(MAIN_PATH)

clean:
	rm -rf $(BUILD_DIR)


