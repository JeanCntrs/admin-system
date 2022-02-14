const socket = new WebSocket('ws://localhost:8000/socket');

window.onload = () => {
    createMenu();
    
    document.getElementById('slcCategory').value = document.getElementById('slcCategory').getAttribute('data-id');
}

const showAlert = () => {
    confirmation().then((result) => {
        if (result.isConfirmed) {
            document.getElementById('frmEditProduct').submit();
            socket.send('editProduct');
        }
    })
}