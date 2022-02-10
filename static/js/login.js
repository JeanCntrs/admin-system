const login = () => {
    const username = document.getElementById('inp_username').value;
    const password = document.getElementById('inp_password').value;

    if (username.trim().length === 0) {
        alert('Username field is required', '', 'error');
        return;
    }

    if (password.trim().length === 0) {
        alert('Password field is required', '', 'error');
        return;
    }

    fetch(`/login/${username}/${password}`)
        .then(response => response.json())
        .then(response => {
            if (response == "1") {
                window.location.href = "/categories"
            } else {
                alert('Username or password is wrong', '', 'error');
                return;
            }
        })
}