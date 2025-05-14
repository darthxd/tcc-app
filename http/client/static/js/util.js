htmx.onLoad(() => {
    const psswd_input = htmx.find("#password")
    const psswd_show_btn = htmx.find("#show_button")
    const psswd_show_icon = psswd_show_btn.children[0]

    htmx.on(psswd_show_btn, "click", () => {
        if(psswd_input.type === "password") {
            psswd_input.type = "text"
            psswd_show_icon.classList.replace("fa-eye", "fa-eye-slash")
        } else {
            psswd_input.type = "password"
            psswd_show_icon.classList.replace("fa-eye-slash", "fa-eye")
        }
    })
})
