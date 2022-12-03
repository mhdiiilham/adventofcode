.PHONY: all build

## Run function will run current year Advent of Code
run:
	go run $(shell date +"%Y")/main.go

## test command will run `go test` command.
test:
	go test -cover ./2022/... ./pkg/input/...

## start command will run `start.sh` script
start:
	./start.sh

bench:
	go test -bench=. \
	./2022/rucksack-reorganization \
	-count 5 \
	-benchmem \
	-run=^#
