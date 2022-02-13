const socket = new WebSocket('ws://localhost:8000/socket');

socket.onopen = () => {
    document.getElementById('lbl_ws_status').innerHTML = 'Connected';
}

socket.onclose = () => {
    document.getElementById('lbl_ws_status').innerHTML = 'Disconnected';
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
            const frm = document.getElementById('frm');
            frm.action = `/products/delete/${element.id}`;
            frm.submit();
        }
    })
}