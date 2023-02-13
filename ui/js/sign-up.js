const form = document.getElementById("form")
const submitButton = document.getElementById("submitTrainerButton")

submitButton.addEventListener("click", function(event){
    event.preventDefault()

    const formData = new FormData(form)
    const signUpData = {}
    console.log(formData)
    for(const [key,value] of formData.entries){
        signUpData.key = value
    }
    console.log(signUpData)
    fetch("http://127.0.0.1:8888/auth/sign-up/trainer",{
        method:"POST",
        headers: {
            "Content-type": "applicatons/json"
        },
        body: JSON.stringify(signUpData)
    })
})
