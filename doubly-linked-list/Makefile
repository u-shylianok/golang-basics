.PHONY: build start run clean tidy

TASK_NAME ?= doubly-linked-list

build:
	go build -o bin/$(TASK_NAME) main.go

start: build
	./bin/$(TASK_NAME)

run:
	go run main.go

clean:
	rm bin/$(TASK_NAME)

tidy:
	go mod tidy