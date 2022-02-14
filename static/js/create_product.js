const socket = new WebSocket('ws://localhost:8000/socket');

window.onload = () => {
    createMenu();
    
    const categoryId = document.getElementById('slcCategory').getAttribute('data-id');
    
    if (categoryId != '0') {
        document.getElementById('slcCategory').value = categoryId;
    }
}

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
            socket.send('createProduct');
        }
    })
}