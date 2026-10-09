# Nazwa pliku binarnego i modułu
BINARY_NAME=dhr

# Parametry kompilacji
BUILD_DIR=bin

.PHONY: all build run clean test install cross-build help

# Domyślne zadanie: buduje aplikację
all: build

## build: Kompiluje plik binarny dhr w katalogu bin/
build:
	@echo "==> Budowanie pliku binarnego $(BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(BINARY_NAME) main.go
	@echo "==> Gotowe: $(BUILD_DIR)/$(BINARY_NAME)"

## run: Uruchamia program bezpośrednio przez go run
run:
	go run main.go

## test: Uruchamia testy jednostkowe
test:
	@echo "==> Uruchamianie testów..."
	go test -v ./...

## clean: Usuwa zbudowane pliki binarne
clean:
	@echo "==> Czyszczenie katalogu $(BUILD_DIR)..."
	rm -rf $(BUILD_DIR)

## install: Instaluje dhr w systemie ($GOPATH/bin)
install:
	@echo "==> Instalowanie $(BINARY_NAME) w $(shell go env GOPATH)/bin..."
	go install .

## cross-build: Kompiluje aplikację na Linux, macOS (Intel/M1) oraz Windows
cross-build: clean
	@echo "==> Kompilacja krzyżowa na różne platformy..."
	@mkdir -p $(BUILD_DIR)
	# Linux 64-bit
	GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o $(BUILD_DIR)/$(BINARY_NAME)-linux-amd64 main.go
	# macOS Intel
	GOOS=darwin GOARCH=amd64 go build -ldflags="-w -s" -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-amd64 main.go
	# macOS Apple Silicon (M1/M2/M3)
	GOOS=darwin GOARCH=arm64 go build -ldflags="-w -s" -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-arm64 main.go
	# Windows 64-bit
	GOOS=windows GOARCH=amd64 go build -ldflags="-w -s" -o $(BUILD_DIR)/$(BINARY_NAME)-windows-amd64.exe main.go
	@echo "==> Pliki binarne utworzone w katalogu $(BUILD_DIR)/"

## help: Wyświetla listę dostępnych komend w Makefile
help:
	@echo "Dostępne komendy Makefile dla $(BINARY_NAME):"
	@sed -n 's/^##//p' $(MAKEFILE_LIST) | column -t -s ':' |  sed -e 's/^/ /'