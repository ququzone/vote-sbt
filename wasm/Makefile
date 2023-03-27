build:
	@go mod tidy
	@tinygo build -o vote-sbt.wasm -scheduler=none --no-debug -target=wasi main.go
	@echo built to vote-sbt.wasm
