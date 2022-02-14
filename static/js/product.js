const socket = new WebSocket('ws://localhost:8000/socket');

socket.onopen = () => {
    document.getElementById('lbl_ws_status').innerHTML = 'Connected';
}

socket.onclose = () => {
    document.getElementById('lbl_ws_status').innerHTML = 'Disconnected';
}

socket.onmessage = (event) => {
    const data = event.data;

    if (data == 'editProduct' || data == 'createProduct') {
        setTimeout(() => {
            document.getElementById('frm_product').submit();
        }, 1000);
    } else if (data.includes('deleteProduct')) {
        const arr = data.split('*');
        const id = arr[1];

        const frm = document.getElementById('frm');
        frm.action = `/products/delete/${id}`;
        frm.submit();
    }
}

window.onload = () => {
    createMenu();

    paginate('table');

    const categoryId = document.getElementById("categoryIdSearched").value;
    document.getElementById("categoryId").value = categoryId;
}

const showDeleteModal = (element) => {
    document.getElementById('txtId').value = element.id;

    confirmation().then((result) => {
        console.log('result', result);
        if (result.isConfirmed) {
            socket.send(`deleteProduct*${element.id}`);
        }
    })
}