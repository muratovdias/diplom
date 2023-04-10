const submit = document.getElementById("submitBtn")     
    
    //request
    submit.addEventListener("click", function(event){
        event.preventDefault()

        const formData = new FormData(document.getElementById("form"));
        const signUpData = {}
        for(const [key,value] of formData){
            signUpData[key] = value
        }
        console.log(JSON.stringify(signUpData))
        fetch("http://127.0.0.1:8888/auth/sign-up",{
            method:"POST",
            headers: {
                "Content-type": "applicatons/json"
            },
            body: JSON.stringify(signUpData)
        })
        .then(response => response.json())
        .then(data => {
            if (data.status !== 200) {
                throw new Error(data.message)
            }
            console.log('Success:', data)
        })
        .catch(error => alert(error));
    })