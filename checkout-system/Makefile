.PHONY: build start run clean tidy

TASK_NAME ?= checkout-system
PRODUCTS_PATH ?= products.json

build:
	go build -o bin/$(TASK_NAME) cmd/$(TASK_NAME)/*.go

start: build
	PRODUCTS_PATH=$(PRODUCTS_PATH) ./bin/$(TASK_NAME)

start-verbose: build
	PRODUCTS_PATH=$(PRODUCTS_PATH) ./bin/$(TASK_NAME) -v

run:
	go run cmd/*.go

clean:
	rm bin/$(TASK_NAME)

tidy:
	go mod tidy