const showAlert = () => {
    const categoryName = document.getElementById('txtName').value;
    const categoryDescription = document.getElementById('txtDescription').value;

    if (categoryName.trim().length === 0) {
        alert('Name field is required', '', 'error');
        return;
    }

    if (categoryDescription.trim().length === 0) {
        alert('Description field is required', '', 'error');
        return;
    }

    confirmation().then((result) => {
        if (result.isConfirmed) {
            document.getElementById('frmCreateCategory').submit();
        }
    })
}