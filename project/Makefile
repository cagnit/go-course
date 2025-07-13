# Путь к директории с main.go
APP_DIR := ./cmd/app

# Имя итогового бинарника
APP_BIN := ../bin/app

# Команда запуска
run:
	go run $(APP_DIR)

# Команда сборки
build:
	go build -o $(APP_BIN) $(APP_DIR)

# Проверка форматирования и статики
link:
	go fmt ./...
	go vet ./...
