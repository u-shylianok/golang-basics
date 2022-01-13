.PHONY: build start run clean tidy

TASK_NAME ?= bst-extended

build:
	go build -o bin/$(TASK_NAME) cmd/$(TASK_NAME)/main.go

start: build
	./bin/$(TASK_NAME)

run:
	go run main.go

clean:
	rm bin/$(TASK_NAME)

tidy:
	go mod tidy