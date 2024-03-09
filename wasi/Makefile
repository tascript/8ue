build:
	@GOOS=wasip1 GOARCH=wasm go build -o main.wasi.wasm main.go

exec: 
	@wasmtime main.wasi.wasm

all: build exec