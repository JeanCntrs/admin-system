const showAlert = () => {
    const productName = document.getElementById('txtProductName').value;
    const productDescription = document.getElementById('txtProductDescription').value;
    const productPrice = document.getElementById('txtProductPrice').value;
    const productStock = document.getElementById('txtProductStock').value;
    const category = document.getElementById('slcCategory').value;

    if (productName.trim().length === 0) {
        alert('Name field is required', '', 'error');
        return;
    }

    if (productDescription.trim().length === 0) {
        alert('Description field is required', '', 'error');
        return;
    }
    
    if (productPrice.trim().length === 0) {
        alert('Price field is required', '', 'error');
        return;
    }

    if (productStock.trim().length === 0) {
        alert('Stock field is required', '', 'error');
        return;
    }

    if (category.trim().length === 0) {
        alert('Category field is required', '', 'error');
        return;
    }

    confirmation().then((result) => {
        if (result.isConfirmed) {
            document.getElementById('frmCreateProduct').submit();
        }
    })
}