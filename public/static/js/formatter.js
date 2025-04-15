rm.addEventListener('input', function () {
    })

htmx.onLoad(() => {
    const rm = htmx.find("#rm")
    htmx.on(rm, "input", () => {
        // Remove tudo que não for número
        let valor = rm.value.replace(/\D/g, '');

        // Limita a 5 caracteres
        if (valor.length > 5) {
          valor = valor.slice(0, 5);
        }

        rm.value = valor;
    })
})
