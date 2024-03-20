copy:
	@cp $(shell go env GOROOT)/misc/wasm/wasm_exec.js ./assets

build:
	@GOOS=js GOARCH=wasm go build -o ./assets/main.wasm ./base/main.go

exec: 
	@go run server.go

all: copy build exec