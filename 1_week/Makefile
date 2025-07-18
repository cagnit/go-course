APP_NAME := app
CMD_PATH := ./cmd/app
BIN_PATH := ../bin

.PHONY: run build link

run:
	go run $(CMD_PATH)/main.go

build:
	@mkdir -p $(BIN_PATH)
	go build -o $(BIN_PATH)/$(APP_NAME) $(CMD_PATH)/main.go

link:
	go fmt ./...
	go vet ./...
