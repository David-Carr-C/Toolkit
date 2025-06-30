GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
ARGS=$(filter-out $@,$(MAKECMDGOALS))

BINARY_NAME=toolkit
MAIN_PACKAGE=./cmd/toolkit  # Punto clave: Especificar paquete principal

BOLD=\033[1m
RED=\033[31m
RESET=\033[0m
YELLOW=\033[33m
MAGENTA=\033[35m
CIAN=\033[36m
WHITE=\033[37m
ARROW=\033[1;34m==>\033[0m 

# Make production build
all: clean build

# Clear go.mod, go.sum, cache and the binary
clean:
	@echo "$(ARROW)$(BOLD)$(YELLOW)Clean started$(RESET)"
	$(GOCLEAN)
	rm -f build/$(BINARY_NAME)
	@echo "$(ARROW)$(BOLD)$(YELLOW)Clean complete$(RESET)"

# Build the binary
build:
	@echo "$(ARROW)$(BOLD)$(RED)Build started$(RESET)"
	$(GOBUILD) -o build/$(BINARY_NAME) -v $(MAIN_PACKAGE)
	@echo "$(ARROW)$(BOLD)$(RED)Build complete$(RESET)"

# Build the binary and run it with the given arguments
run:
	@echo "$(ARROW)$(BOLD)$(RED)Build started$(RESET)"
	$(GOBUILD) -o build/$(BINARY_NAME) -v $(MAIN_PACKAGE)
	@echo "$(ARROW)$(BOLD)$(RED)Build complete$(RESET)"
	@echo "$(ARROW)$(BOLD)$(MAGENTA)Run started$(RESET)"
	./build/$(BINARY_NAME) $(ARGS)
	@echo "Run complete"
	@echo "$(ARROW)$(BOLD)$(MAGENTA)Run complete$(RESET)"
	@exit 1


# Run the tests
test:
	@echo "$(ARROW)$(BOLD)$(CIAN)Test started$(RESET)"
	$(GOTEST) -v ./...
	@echo "$(ARROW)$(BOLD)$(CIAN)Test complete$(RESET)"

# Proof whatever you want in Makefile
proof:
	@echo "$(ARROW)$(BOLD)$(WHITE)Proof started$(RESET)"
	./build/$(BINARY_NAME) --help
	@echo "---------------------------------------------------------"
	./build/$(BINARY_NAME) $(ARGS) --help
	@echo "${ARROW}$(BOLD)$(WHITE)Proof complete$(RESET)"

.PHONY: all clean build run test proof
