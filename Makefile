MAIN := ./cmd/lambda/main.go

run:
	go run $(MAIN) $(args)

build:
	go build $(MAIN)

.PHONY: run