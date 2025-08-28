# Makefile for Zend
# ========================

# Paths
FRONTEND_DIR := web
CONFIG_DIR := ${XDG_CONFIG_HOME:-$(HOME)/.config}/zend
CONFIG_FILE := $(CONFIG_DIR)/config.yaml
CLI_BIN := zend

# Default task
.PHONY: all
all: build

# -------------------------
# Cli
# -------------------------
.PHONY: cli
cli:
	@echo "Building cli..."
	go mod tidy
	go build -o zend

.PHONY: cli-run
cli-run: cli
	@echo "Running cli..."
	$(CLI_BIN)

# -------------------------
# Frontend
# -------------------------
.PHONY: frontend
frontend:
	@echo "Installing frontend dependencies..."
	cd $(FRONTEND_DIR) && npm install
	@echo "Building frontend..."
	cd $(FRONTEND_DIR) && npm run build

.PHONY: frontend-dev
frontend-dev:
	cd $(FRONTEND_DIR) && npm run dev

# -------------------------
# Config
# -------------------------
.PHONY: config
config:
	@if [ ! -d "$(CONFIG_DIR)" ]; then mkdir -p "$(CONFIG_DIR)"; fi
	@if [ ! -f "$(CONFIG_FILE)" ]; then cp $ config/default.yaml $(CONFIG_FILE); fi
	@echo "Config ready at $(CONFIG_FILE)"

# -------------------------
# Clean
# -------------------------
.PHONY: clean
clean:
	@echo "Cleaning cli binary and frontend build..."
	rm -f $(CLI_BIN)
	rm -rf $(FRONTEND_DIR)/dist

# -------------------------
# Full build
# -------------------------
.PHONY: build
build: config cli frontend
	@echo "Zend build completed."

# -------------------------
# Dev
# -------------------------
.PHONY: dev
dev:
	@echo "Starting cli and frontend dev servers..."
	@$(MAKE) cli-run &
	cd $(FRONTEND_DIR) && npm run dev
