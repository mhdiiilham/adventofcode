.PHONY: all build

# Run function will run current year Advent of Code
run:
	go run $(shell date +"%Y")/main.go

test:
	go test -cover -v ./2022/... ./pkg/input/...

start:
	./start.sh

