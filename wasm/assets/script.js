const goWasm = new Go()

WebAssembly.instantiateStreaming(fetch("monkey.wasm"), goWasm.importObject)
    .then((result)=>{
        goWasm.run(result.instance)       

        document.getElementById("get-html").addEventListener("click", () => {
            document.body.innerHTML = getHtml()
        })
    })
    // https://youtu.be/10Mz3z-W1BE