.PHONY: build start run clean

TASK_NAME ?= task01

build:
	go build -o bin/$(TASK_NAME) cmd/*.go

start:
	make build
	./bin/$(TASK_NAME)

run:
	go run cmd/*.go

clean:
	rm bin/$(TASK_NAME)