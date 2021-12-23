window.onload = () => {
    paginate('table');
}

const showDeleteModal = (categoryId) => {
    document.getElementById('txtId').value = categoryId;

    confirmation().then((result) => {
        if (result.isConfirmed) {
            const frm = document.getElementById('frm');
            frm.action = `/categories/delete/${categoryId}`;
            frm.submit();
        }
    })
}