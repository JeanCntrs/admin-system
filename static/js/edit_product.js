window.onload = () => {
    document.getElementById('slcCategory').value = document.getElementById('slcCategory').getAttribute('data-id');
}

const showAlert = () => {
    confirmation().then((result) => {
        if (result.isConfirmed) {
            document.getElementById('frmEditProduct').submit();
        }
    })
}