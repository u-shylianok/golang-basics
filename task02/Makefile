.PHONY: build start run clean tidy start-env-all start-args-all start-env-args-override start-env-args-parts start-env-args-parts-override

TASK_NAME ?= task02
INPUT_PATH ?= XMLConfig.xml
OUTPUT_PATH ?= JSONConfig.json

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

start-env-all: build
	INPUT_PATH=$(INPUT_PATH) OUTPUT_PATH=$(OUTPUT_PATH) ./bin/$(TASK_NAME)

start-args-all: build
	./bin/$(TASK_NAME) $(INPUT_PATH) $(OUTPUT_PATH)

start-env-args-override: build
	INPUT_PATH=$(INPUT_PATH) OUTPUT_PATH=$(OUTPUT_PATH) ./bin/$(TASK_NAME) $(INPUT_PATH) arg_$(OUTPUT_PATH)

start-env-args-parts: build
	OUTPUT_PATH=$(OUTPUT_PATH) ./bin/$(TASK_NAME) $(INPUT_PATH)

start-env-args-parts-override: build
	INPUT_PATH=$(INPUT_PATH) OUTPUT_PATH=env_$(OUTPUT_PATH) ./bin/$(TASK_NAME) $(INPUT_PATH)
