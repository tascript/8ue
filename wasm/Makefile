copy:
	@cp $(shell go env GOROOT)/misc/wasm/wasm_exec.js .

build:
	@GOOS=js GOARCH=wasm go build -o main.wasm main.go