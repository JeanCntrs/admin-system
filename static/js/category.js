const socket = new WebSocket('ws://localhost:8000/socket');

socket.onopen = () => {
    document.getElementById('lbl_ws_status').innerHTML = 'Connected';
}

socket.onclose = () => {
    document.getElementById('lbl_ws_status').innerHTML = 'Disconnected';
}

socket.onmessage = (event) => {
    const data = event.data;
    
    if (data.includes('deleteCategory')) {
        const arr = data.split('*');
        const id = arr[1];

        const frm = document.getElementById('frm');
        frm.action = `/categories/delete/${id}`;
        frm.submit();
    } else if (data == 'editCategory' || data == 'createCategory') {
        setTimeout(() => {
            // document.getElementById('frm_category').submit();
            document.location.reload();
        }, 1000);
    }
}

window.onload = () => {
    createMenu()
    paginate('table');
}

const showDeleteModal = (element) => {
    document.getElementById('txtId').value = element.id;

    confirmation().then((result) => {
        if (result.isConfirmed) {
            // const frm = document.getElementById('frm');
            // frm.action = `/categories/delete/${element.id}`;
            // frm.submit();
            socket.send(`deleteCategory*${element.id}`);
        }
    })
}