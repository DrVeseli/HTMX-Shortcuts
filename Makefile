# Makefile

# Define variables
TAILWIND_CMD = ./tailwindcss
INPUT_FILE = ./templates/input.css
OUTPUT_FILE = ./templates/output.css
AIR_CMD = air -c .air.toml

# Target to start Tailwind CSS in watch mode
.PHONY: tailwind
tailwind:
	@echo "Starting Tailwind CSS watch..."
	$(TAILWIND_CMD) -i $(INPUT_FILE) -o $(OUTPUT_FILE) --watch

# Target to start air
.PHONY: air
air:
	@echo "Starting air..."
	$(AIR_CMD)

# Target to start both Tailwind CSS and air concurrently
.PHONY: all
all: tailwind air

# Clean target to remove generated files (if needed)
.PHONY: clean
clean:
	rm -f $(OUTPUT_FILE)

# Help target
.PHONY: help
help:
	@echo "Makefile for Tailwind CSS and air"
	@echo "Usage:"
	@echo "  make tailwind - Start Tailwind CSS in watch mode"
	@echo "  make air      - Start air"
	@echo "  make all      - Start both Tailwind CSS and air concurrently"
	@echo "  make clean    - Remove generated output files"
	@echo "  make help     - Show this help message"
