BIN_NAME=forge
APP_NAME=forge
CMD_DIR=./cmd
BIN_DIR=./bin

.PHONY: build run clean comm

build:
	@mkdir -p $(BIN_DIR)		#create bin dir if it doesn't exist
	go build -o	 $(BIN_DIR)/$(BIN_NAME) $(CMD_DIR)

run:
	@go run $(CMD_DIR)

clean:
	rm -f $(BIN_DIR)/$(BIN_NAME)

comm:
	@git add .
	@git commit -m "$(msg)"
	@git push