GOPATH:=$(shell go env GOPATH)

.PHONY: run test

run:
	go run main.go

test:
	@go test -v -race ./... | sed ''/PASS/s//$$(printf "\033[32mPASS\033[0m")/'' | sed ''/FAIL/s//$$(printf "\033[31mFAIL\033[0m")/''