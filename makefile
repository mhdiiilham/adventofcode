.PHONY: all build

## Run function will run current year Advent of Code
run:
	go run $(shell date +"%Y")/main.go

## test command will run `go test` command.
# but only on current year directory
test:
	go test -cover ./2022/... ./pkg/input/...

## start command will run `start.sh` script
start:
	./start.sh
