window.onload = () => {
    paginate('table');
}

const showAlert = () => {
    confirmation().then((result) => {
        if (result.isConfirmed) {
            alert()
        }
    })
}