# Makefile

# Variables
GO_CMD = go
SRC_FILE = vorto/main.go
TARGET = vorto
BIN_DIR = ./bin
TEST301 = "./training/Problems/problem1.txt"
TEST302 = "./training/Problems/problem2.txt"
TEST303 = "./training/Problems/problem3.txt"
TEST304 = "./training/Problems/problem4.txt"
TEST305 = "./training/Problems/problem5.txt"
TEST306 = "./training/Problems/problem6.txt"
TEST307 = "./training/Problems/problem7.txt"
TEST308 = "./training/Problems/problem8.txt"
TEST309 = "./training/Problems/problem9.txt"
TEST310 = "./training/Problems/problem10.txt"
TEST311 = "./training/Problems/problem11.txt"
TEST312 = "./training/Problems/problem12.txt"
TEST313 = "./training/Problems/problem13.txt"
TEST314 = "./training/Problems/problem14.txt"
TEST315 = "./training/Problems/problem15.txt"
TEST316 = "./training/Problems/problem16.txt"
TEST317 = "./training/Problems/problem17.txt"
TEST318 = "./training/Problems/problem18.txt"
TEST319 = "./training/Problems/problem19.txt"
TEST320 = "./training/Problems/problem20.txt"





# Default target
.PHONY: all
all: build

# Build target
.PHONY: build
build:
	@if [ ! -d "$(BIN_DIR)" ]; then \
		echo "Directory $(BIN_DIR) does not exist, creating it."; \
		mkdir -p $(BIN_DIR); \
		chmod 755 $(BIN_DIR); \
	fi
	@echo "Building $(TARGET) into $(BIN_DIR)"; \
	$(GO_CMD) build -o $(BIN_DIR)/$(TARGET) $(SRC_FILE)

# Run target
.PHONY: run
run:
	$(GO_CMD) run $(SRC_FILE) v100 -T ./training/Problems/problem20.txt

# Clean target
.PHONY: clean
clean:
	@echo "Cleaning up..."
	@rm -f $(BIN_DIR)/$(TARGET)
	@echo "Done."

# Test target
.PHONY: vorto_test
vorto_test: build
	@echo "Running several versions of the Problems..."
	@echo "<<<Test 100 START>>>"
	@$(BIN_DIR)/$(TARGET) -h
	@echo "<<<Test 100 COMPLETE >>>"
	@echo "<<<Test 200 START"
	@$(BIN_DIR)/$(TARGET) -h v100 
	@echo "<<<Test 200 COMPLETE>>>"
	@echo "<<<Test 301 START >>>"
	@$(BIN_DIR)/$(TARGET) v100 -T $(TEST301)
	@echo "<<<Test 301 COMPLETE>>>"
	@echo "<<<Test 302 START >>>"
	@$(BIN_DIR)/$(TARGET) v100 -T $(TEST302)
	@echo "<<<Test 302 COMPLETE>>>"
	@echo "<<<Test 303 START >>>"
	@$(BIN_DIR)/$(TARGET) v100 -T $(TEST303)
	@echo "<<<Test 303 COMPLETE>>>"
	@echo "<<<Test 304 START >>>"
	@$(BIN_DIR)/$(TARGET) v100 -T $(TEST304)
	@echo "<<<Test 304 COMPLETE>>>"
	@echo "<<<Test 305 START >>>"
	@$(BIN_DIR)/$(TARGET) v100 -T $(TEST305)
	@echo "<<<Test 305 COMPLETE>>>"
	@echo "<<<Test 306 START >>>"
	@$(BIN_DIR)/$(TARGET) v100 -T $(TEST306)
	@echo "<<<Test 306 COMPLETE>>>"
	@echo "<<<Test 307 START >>>"
	@$(BIN_DIR)/$(TARGET) v100 -T $(TEST307)
	@echo "<<<Test 307 COMPLETE>>>"
	@echo "<<<Test 308 START >>>"
	@$(BIN_DIR)/$(TARGET) v100 -T $(TEST308)
	@echo "<<<Test 308 COMPLETE>>>"
	@echo "<<<Test 309 START >>>"
	@$(BIN_DIR)/$(TARGET) v100 -T $(TEST309)
	@echo "<<<Test 309 COMPLETE>>>"
	@echo "<<<Test 310 START >>>"
	@$(BIN_DIR)/$(TARGET) v100 -T $(TEST310)
	@echo "<<<Test 310 COMPLETE>>>"
	@echo "<<<Test 311 START >>>"
	@$(BIN_DIR)/$(TARGET) v100 -T $(TEST311)
	@echo "<<<Test 311 COMPLETE>>>"
	@echo "<<<Test 312 START >>>"
	@$(BIN_DIR)/$(TARGET) v100 -T $(TEST312)
	@echo "<<<Test 312 COMPLETE>>>"
	@echo "<<<Test 313 START >>>"
	@$(BIN_DIR)/$(TARGET) v100 -T $(TEST313)
	@echo "<<<Test 313 COMPLETE>>>"
	@echo "<<<Test 314 START >>>"
	@$(BIN_DIR)/$(TARGET) v100 -T $(TEST314)
	@echo "<<<Test 314 COMPLETE>>>"
	@echo "<<<Test 315 START >>>"
	@$(BIN_DIR)/$(TARGET) v100 -T $(TEST315)
	@echo "<<<Test 315 COMPLETE>>>"
	@echo "<<<Test 316 START >>>"
	@$(BIN_DIR)/$(TARGET) v100 -T $(TEST316)
	@echo "<<<Test 316 COMPLETE>>>"
	@echo "<<<Test 317 START >>>"
	@$(BIN_DIR)/$(TARGET) v100 -T $(TEST317)
	@echo "<<<Test 317 COMPLETE>>>"
	@echo "<<<Test 318 START >>>"
	@$(BIN_DIR)/$(TARGET) v100 -T $(TEST318)
	@echo "<<<Test 318 COMPLETE>>>"
	@echo "<<<Test 319 START >>>"
	@$(BIN_DIR)/$(TARGET) v100 -T $(TEST319)
	@echo "<<<Test 319 COMPLETE>>>"
	@echo "<<<Test 320 START >>>"
	@$(BIN_DIR)/$(TARGET) v100 -T $(TEST320)
	@echo "<<<Test 320 COMPLETE>>>"



# Help target
.PHONY: help
help:
	@echo "Usage:"
	@echo "  make           - Build the executable"
	@echo "  make build     - Build the executable"
	@echo "  make run       - Run the Go program"
	@echo "  make clean     - Clean up"
	@echo "  make vorto_test- Run several versions of the Problems"
	@echo "  make help      - Show this help message"
