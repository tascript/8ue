build:
	@GOOS=wasip1 GOARCH=wasm go build -o main.wasi.wasm main.go

install: 
	@go install github.com/stealthrocket/wasi-go/cmd/wasirun@latest

exec: 
	@wasirun main.wasi.wasm

all: build install exec