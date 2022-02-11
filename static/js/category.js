window.onload = () => {
    createMenu()
    paginate('table');
}

const showDeleteModal = (element) => {
    document.getElementById('txtId').value = element.id;

    confirmation().then((result) => {
        if (result.isConfirmed) {
            const frm = document.getElementById('frm');
            frm.action = `/categories/delete/${element.id}`;
            frm.submit();
        }
    })
}