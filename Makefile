#!/use/bin/env make

.PHONY: build
build:
	@ mkdir -p bin
	go build -o ./bin/github-watch ./cmd/github-watch/

.PHONY: run
run: build
	@./bin/github-watch
