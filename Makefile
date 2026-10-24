BINARY_NAME=dhr

BUILD_DIR=bin

.PHONY: all build run clean test install cross-build help

all: build

build:
	@echo "==> Budowanie pliku binarnego $(BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(BINARY_NAME) main.go
	@echo "==> Gotowe: $(BUILD_DIR)/$(BINARY_NAME)"

run:
	go run main.go

test:
	@echo "==> Uruchamianie testów..."
	go test -v ./...

clean:
	@echo "==> Czyszczenie katalogu $(BUILD_DIR)..."
	rm -rf $(BUILD_DIR)

install:
	@echo "==> Instalowanie $(BINARY_NAME) w $(shell go env GOPATH)/bin..."
	go install .

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