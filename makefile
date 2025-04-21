BINARY_NAME=main.out
APP_NAME=forge
CMD_DIR=./cmd

.PHONY: run clean build

build:
go run $(CMD_DIR)

run:
go build -o $(APP_NAME) $(CMD_DIR)

clean:
rm -f forge forge.exe