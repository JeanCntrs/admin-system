const showAlert = () => {
    confirmation().then((result) => {
        if (result.isConfirmed) {
            alert()
            document.getElementById('frm_create_category').submit();
        }
    })
}