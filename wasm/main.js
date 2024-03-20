const go = new Go()
const { instance } = await WebAssembly.instantiateStreaming(
  fetch("main.wasm"),
  go.importObject
)
const greet = () => goGreet() 
