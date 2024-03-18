build:
	@GOOS=wasip1 GOARCH=wasm go build -o main.wasi.wasm main.go

install:
	@brew install wasmtime

exec: 
	@wasmtime main.wasi.wasm

all: build exec