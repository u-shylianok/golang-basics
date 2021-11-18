.PHONY: build start run clean tidy

TASK_NAME ?= checkout-system

build:
	go build -o bin/$(TASK_NAME) cmd/$(TASK_NAME)/*.go

start: build
	./bin/$(TASK_NAME)

run:
	go run cmd/*.go

clean:
	rm bin/$(TASK_NAME)

tidy:
	go mod tidy