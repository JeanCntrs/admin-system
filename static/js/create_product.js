const showAlert = () => {
    confirmation().then((result) => {
        if (result.isConfirmed) {
            document.getElementById('frmCreateProduct').submit();
        }
    })
}