(async () => {
  const go = new Go()
  const { instance } = await WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject)
  go.run(instance)
})()
const greet = () => goGreet()
