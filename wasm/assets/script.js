const goWasm = new Go()

WebAssembly.instantiateStreaming(fetch("monkey.wasm"), goWasm.importObject)
    .then((result)=>{
        goWasm.run(result.instance)

        document.getElementById("interpretBtn").addEventListener("click", () => {
            let inputText = document.getElementById("interpretInput").value;
            monkeyInterpret(inputText);
            // document.body.innerHTML = getHtml("Hello World")
        })
    })
    // https://youtu.be/10Mz3z-W1BE