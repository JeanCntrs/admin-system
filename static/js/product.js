window.onload = () => {
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