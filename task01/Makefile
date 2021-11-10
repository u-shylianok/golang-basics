.PHONY: build start run clean tidy

TASK_NAME ?= task01

build:
	go build -o bin/$(TASK_NAME) cmd/*.go

start: build
	./bin/$(TASK_NAME)

run:
	go run cmd/*.go

clean:
	rm bin/$(TASK_NAME)

tidy:
	go mod tidy