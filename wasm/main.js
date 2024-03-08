import './wasm_exec.js'

const path = new URL("main.wasm", import.meta.url)
const go = new Go()
const { instance } = await WebAssembly.instantiateStreaming(
  fetch(path),
  {
    env: {
      add: function add(v) {
        let res = 0
        for (let i = 1; i <= v; i++) {
          res += i
        }
        return res
      }
    },
    ...go.importObject
  }
)
go.run(instance)
